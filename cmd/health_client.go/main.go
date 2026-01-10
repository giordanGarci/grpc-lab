package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
)

func main() {
	fmt.Println("starting health client")

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	md := metadata.Pairs("authorization", "your-auth-token")
	ctxWithMeta := metadata.NewOutgoingContext(context.Background(), md)

	ctx, cancel := context.WithTimeout(ctxWithMeta, time.Second*10)
	defer cancel()

	healthClient := healthpb.NewHealthClient(conn)

	resp, err := healthClient.Check(ctx, &healthpb.HealthCheckRequest{
		Service: "",
	})

	if err != nil {
		log.Fatalf("Health check failed: %v", err)
	}

	fmt.Printf("Server Status: %s\n", resp.GetStatus())

}
