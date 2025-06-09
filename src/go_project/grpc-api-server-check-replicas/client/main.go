package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "go_grpc/proto"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewReplicaServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	resp, err := client.GetDeploymentReplicas(ctx, &pb.ReplicaRequest{
		Namespace: "default",
		Name:      "nginx-deployment",
	})
	if err != nil {
		log.Fatalf("Error calling GetDeploymentReplicas: %v", err)
	}

	fmt.Println("Name", resp.Name)
	fmt.Println("Available Replicas:", resp.AvailableReplicas)
}
