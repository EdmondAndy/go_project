package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Returns a Kubernetes clientset from in-cluster config or kubeconfig
func getKubeClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		// fallback to kubeconfig
		kubeconfig := os.Getenv("KUBECONFIG")
		if kubeconfig == "" {
			kubeconfig = fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	config.Timeout = 120 * time.Second
	return kubernetes.NewForConfig(config)
}

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	clientset, err := getKubeClient()
	if err != nil {
		http.Error(w, "❌ Cannot create Kubernetes client", http.StatusInternalServerError)
		return
	}

	_, err = clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		http.Error(w, "❌ Cannot reach Kubernetes API", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "✅ Kubernetes connectivity is healthy")
}

func main() {
	http.HandleFunc("/healthz", healthHandler)
	port := "8080"
	fmt.Printf("🔧 Starting health check server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
