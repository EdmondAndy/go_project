package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeClient() (*kubernetes.Clientset, error) {
	// Try in-cluster config first
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fall back to kubeconfig file
		kubeconfig := os.Getenv("KUBECONFIG")
		if kubeconfig == "" {
			home := os.Getenv("HOME")
			kubeconfig = fmt.Sprintf("%s/.kube/config", home)
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	return clientset, err
}

func deploymentsHandler(w http.ResponseWriter, r *http.Request) {
	clientset, err := getKubeClient()
	if err != nil {
		http.Error(w, "Failed to create Kubernetes client", http.StatusInternalServerError)
		return
	}

	deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		http.Error(w, "Failed to list deployments", http.StatusInternalServerError)
		return
	}

	var result []map[string]string
	for _, d := range deployments.Items {
		result = append(result, map[string]string{
			"name":      d.Name,
			"namespace": d.Namespace,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/deployments", deploymentsHandler)
	port := "8080"
	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
