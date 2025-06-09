package test

import (
	"context"
	"testing"
	"time"

	pb "go_grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetDeploymentReplicas(t *testing.T) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewReplicaServiceClient(conn)

	// Wait for deployment to be available
	time.Sleep(5 * time.Second)

	resp, err := client.GetDeploymentReplicas(context.TODO(), &pb.ReplicaRequest{
		Namespace: "default",
		Name:      "test-deployment",
	})
	if err != nil {
		t.Fatalf("Failed to get deployment: %v", err)
	}

	if resp.Replicas != 2 {
		t.Errorf("Expected 2 replicas, got %d", resp.Replicas)
	}
}
