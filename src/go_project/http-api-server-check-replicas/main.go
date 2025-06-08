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

// DeploymentReplicaInfo holds the replica info to return as JSON
type DeploymentReplicaInfo struct {
	Name              string `json:"name"`
	Namespace         string `json:"namespace"`
	DesiredReplicas   int32  `json:"desired_replicas"`
	AvailableReplicas int32  `json:"available_replicas"`
	ReadyReplicas     int32  `json:"ready_replicas"`
}

// GetKubeClient initializes the Kubernetes client
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

// Handler for GET /replicas?namespace=default&name=my-deployment
func GetDeploymentReplicas(w http.ResponseWriter, r *http.Request) {
	namespace := r.URL.Query().Get("namespace")
	name := r.URL.Query().Get("name")

	if namespace == "" || name == "" {
		http.Error(w, "Missing namespace or name parameter", http.StatusBadRequest)
		return
	}

	clientset, err := GetKubeClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating k8s client: %v", err), http.StatusInternalServerError)
		return
	}

	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	deployment, err := deploymentsClient.Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting deployment: %v", err), http.StatusInternalServerError)
		return
	}

	info := DeploymentReplicaInfo{
		Name:              deployment.Name,
		Namespace:         deployment.Namespace,
		DesiredReplicas:   *deployment.Spec.Replicas,
		AvailableReplicas: deployment.Status.AvailableReplicas,
		ReadyReplicas:     deployment.Status.ReadyReplicas,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func main() {
	http.HandleFunc("/replicas", GetDeploymentReplicas)

	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
