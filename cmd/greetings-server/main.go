package main

import (
	"fmt"
	"net"

	"grpc-lab/internal/greetings"
	pb "grpc-lab/pb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting server grpc")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	fmt.Println("server grpc started on port 50051")

	service := greetings.NewService()
	handler := greetings.NewHandler(service)

	pb.RegisterHelloServiceServer(grpcServer, handler)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
