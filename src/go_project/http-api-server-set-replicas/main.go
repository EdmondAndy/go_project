package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type ScaleRequest struct {
	Namespace  string `json:"namespace"`
	Deployment string `json:"deployment"`
	Replicas   int32  `json:"replicas"`
}

func GetKubeClient() (*kubernetes.Clientset, error) {
	// Try in-cluster config first
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fall back to local kubeconfig
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	return clientset, err
}

func main() {
	// Setup in-cluster Kubernetes config
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	log.Fatalf("Error getting in-cluster config: %v", err)
	// }

	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	log.Fatalf("Error creating Kubernetes client: %v", err)
	// }

	http.HandleFunc("/scale", func(w http.ResponseWriter, r *http.Request) {
		clientset, err := GetKubeClient()

		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating k8s client: %v", err), http.StatusInternalServerError)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var req ScaleRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Fetch the current Deployment
		deployClient := clientset.AppsV1().Deployments(req.Namespace)
		deployment, err := deployClient.Get(context.TODO(), req.Deployment, metav1.GetOptions{})
		if err != nil {
			http.Error(w, "Deployment not found", http.StatusNotFound)
			return
		}

		// Set replicas
		deployment.Spec.Replicas = &req.Replicas

		// Update Deployment
		_, err = deployClient.Update(context.TODO(), deployment, metav1.UpdateOptions{})
		if err != nil {
			http.Error(w, "Failed to update deployment", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Deployment %s scaled to %d replicas\n", req.Deployment, req.Replicas)
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
