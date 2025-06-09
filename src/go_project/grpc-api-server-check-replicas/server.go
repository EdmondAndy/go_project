package main

import (
	"context"

	pb "go_grpc/proto"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type server struct {
	pb.UnimplementedReplicaServiceServer
}

func (s *server) GetDeploymentReplicas(ctx context.Context, req *pb.ReplicaRequest) (*pb.ReplicaResponse, error) {
	clientset, err := getKubeClient()
	if err != nil {
		return nil, err
	}

	deploy, err := clientset.AppsV1().Deployments(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return &pb.ReplicaResponse{
		Name:              deploy.Name,
		Namespace:         deploy.Namespace,
		Replicas:          deploy.Status.Replicas,
		ReadyReplicas:     deploy.Status.ReadyReplicas,
		AvailableReplicas: deploy.Status.AvailableReplicas,
		UpdatedReplicas:   deploy.Status.UpdatedReplicas,
	}, nil
}
