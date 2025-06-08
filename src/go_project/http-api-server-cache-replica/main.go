package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

type ReplicaInfo struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Replicas  int32  `json:"replicas"`
	Ready     int32  `json:"ready"`
	Available int32  `json:"available"`
	Updated   int32  `json:"updated"`
}

var (
	replicaCache = make(map[string]ReplicaInfo) // key: namespace/name
	cacheLock    sync.RWMutex
)

func cacheKey(ns, name string) string {
	return fmt.Sprintf("%s/%s", ns, name)
}

func getClientSet() (*kubernetes.Clientset, error) {
	// Use in-cluster config if running inside K8s
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

func startWatcher(clientset *kubernetes.Clientset, stopCh <-chan struct{}) {
	factory := informers.NewSharedInformerFactory(clientset, 0)
	informer := factory.Apps().V1().Deployments().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: updateCache,
		UpdateFunc: func(oldObj, newObj interface{}) {
			updateCache(newObj)
		},
		DeleteFunc: func(obj interface{}) {
			if d, ok := obj.(*appsv1.Deployment); ok {
				cacheLock.Lock()
				delete(replicaCache, cacheKey(d.Namespace, d.Name))
				cacheLock.Unlock()
			}
		},
	})

	go factory.Start(stopCh)

	if !cache.WaitForCacheSync(stopCh, informer.HasSynced) {
		log.Fatal("Failed to sync cache")
	}
	log.Println("Watcher synced and running")
}

func updateCache(obj interface{}) {
	d, ok := obj.(*appsv1.Deployment)
	if !ok {
		return
	}
	info := ReplicaInfo{
		Name:      d.Name,
		Namespace: d.Namespace,
		Replicas:  *d.Spec.Replicas,
		Ready:     d.Status.ReadyReplicas,
		Available: d.Status.AvailableReplicas,
		Updated:   d.Status.UpdatedReplicas,
	}

	cacheLock.Lock()
	defer cacheLock.Unlock()
	replicaCache[cacheKey(d.Namespace, d.Name)] = info
}

func replicasHandler(w http.ResponseWriter, r *http.Request) {
	cacheLock.RLock()
	defer cacheLock.RUnlock()

	results := make([]ReplicaInfo, 0, len(replicaCache))
	for _, v := range replicaCache {
		results = append(results, v)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	clientset, err := getClientSet()
	if err != nil {
		log.Fatalf("Error building clientset: %v", err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)
	defer runtime.HandleCrash()
	startWatcher(clientset, stopCh)

	http.HandleFunc("/replicas", replicasHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
