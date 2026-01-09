package main

import (
	"fmt"
	"net"

	"grpc-lab/internal/age"
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

	greetingsService := greetings.NewService()
	greetingsHandler := greetings.NewHandler(greetingsService)

	ageService := age.NewService()
	ageHandler := age.NewHandler(ageService)

	pb.RegisterHelloServiceServer(grpcServer, greetingsHandler)
	pb.RegisterAgeServiceServer(grpcServer, ageHandler)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
