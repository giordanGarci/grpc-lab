package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"grpc-lab/pb"

	// Este é o import que estava faltando para o 'runtime'
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("starting http server")

	ctx := context.Background()
	mux := runtime.NewServeMux()

	// Registra o handler do gateway apontando para onde seu gRPC está rodando
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	log.Println("HTTP Gateway rodando na porta :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
