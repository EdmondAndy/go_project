package main

import (
	"context"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	replicaspb "replicaspb" // adjust import path if needed
)

type server struct {
	replicaspb.UnimplementedReplicaServiceServer
	client *kubernetes.Clientset
}

func (s *server) GetReplicaCount(ctx context.Context, req *replicaspb.GetReplicaRequest) (*replicaspb.GetReplicaResponse, error) {
	deploy, err := s.client.AppsV1().Deployments(req.Namespace).Get(ctx, req.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	var specReplicas int32 = 0
	if deploy.Spec.Replicas != nil {
		specReplicas = *deploy.Spec.Replicas
	}

	return &replicaspb.GetReplicaResponse{
		Name:             deploy.Name,
		Namespace:        deploy.Namespace,
		Replicas:         specReplicas,
		ReadyReplicas:    deploy.Status.ReadyReplicas,
		AvailableReplicas: deploy.Status.AvailableReplicas,
		UpdatedReplicas:  deploy.Status.UpdatedReplicas,
	}, nil
}
