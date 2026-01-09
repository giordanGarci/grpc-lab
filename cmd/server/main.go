package main

import (
	"fmt"
	"net"

	"grpc-lab/internal/age"
	"grpc-lab/internal/greetings"
	"grpc-lab/internal/interceptors"
	pb "grpc-lab/pb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting server grpc")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.LoggerInterceptor),
	)

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
