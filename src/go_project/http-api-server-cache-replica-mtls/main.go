package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type ReplicaInfo struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Replicas  int32  `json:"replicas"`
	Ready     int32  `json:"readyReplicas"`
	Available int32  `json:"availableReplicas"`
	Updated   int32  `json:"updatedReplicas"`
}

var (
	cache      = make(map[string]ReplicaInfo)
	cacheMutex = sync.RWMutex{}
)

func cacheKey(ns, name string) string {
	return ns + "/" + name
}

func startWatcher(clientset *kubernetes.Clientset) {
	watcher, err := clientset.AppsV1().Deployments("").Watch(context.TODO(), metav1.ListOptions{
		FieldSelector: fields.Everything().String(),
	})
	if err != nil {
		log.Fatalf("Failed to start watcher: %v", err)
	}

	go func() {
		for event := range watcher.ResultChan() {
			switch event.Type {
			case watch.Added, watch.Modified:
				if d, ok := event.Object.(*appsv1.Deployment); ok {
					info := ReplicaInfo{
						Name:      d.Name,
						Namespace: d.Namespace,
						Replicas:  *d.Spec.Replicas,
						Ready:     d.Status.ReadyReplicas,
						Available: d.Status.AvailableReplicas,
						Updated:   d.Status.UpdatedReplicas,
					}
					cacheMutex.Lock()
					cache[cacheKey(d.Namespace, d.Name)] = info
					cacheMutex.Unlock()
				}
			case watch.Deleted:
				if d, ok := event.Object.(*appsv1.Deployment); ok {
					cacheMutex.Lock()
					delete(cache, cacheKey(d.Namespace, d.Name))
					cacheMutex.Unlock()
				}
			}
		}
		log.Println("Deployment watcher closed. Restarting...")
		startWatcher(clientset)
	}()
}

func replicasHandler(w http.ResponseWriter, r *http.Request) {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()

	var result []ReplicaInfo
	for _, v := range cache {
		result = append(result, v)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func getKubeClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fallback to kubeconfig (for local dev)
		kubeconfig := os.Getenv("KUBECONFIG")
		if kubeconfig == "" {
			kubeconfig = fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	return kubernetes.NewForConfig(config)
}

func loadMTLSServer(certFile, keyFile, caFile string) (*http.Server, error) {
	// Load server certificate and key
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load cert/key: %w", err)
	}

	// Load CA for client cert verification
	caCert, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read CA cert: %w", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		MinVersion:   tls.VersionTLS12,
	}

	srv := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/replicas", replicasHandler)
	return srv, nil
}

func main() {
	clientset, err := getKubeClient()
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	startWatcher(clientset)

	server, err := loadMTLSServer("server-local.crt", "server-local.key", "ca-local.crt")
	if err != nil {
		log.Fatalf("Failed to configure mTLS server: %v", err)
	}

	log.Println("Server running on https://localhost:8443 with mTLS")
	log.Fatal(server.ListenAndServeTLS("", "")) // TLS config provided manually
}
