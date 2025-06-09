package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "go_grpc/proto"
)

func main() {

	conn, _ := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewReplicaServiceClient(conn)
	resp, _ := client.GetDeploymentReplicas(context.TODO(), &pb.ReplicaRequest{
		Namespace: "default",
		Name:      "nginx-deployment",
	})
	fmt.Println("Available Replicas:", resp.AvailableReplicas)

}
