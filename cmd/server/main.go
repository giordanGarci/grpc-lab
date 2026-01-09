package main

import (
	"fmt"
	"net"

	"grpc-lab/internal/age"
	"grpc-lab/internal/greetings"
	"grpc-lab/internal/interceptors"
	"grpc-lab/internal/slow"
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

	slowService := slow.NewService()
	slowHandler := slow.NewHandler(slowService)

	pb.RegisterHelloServiceServer(grpcServer, greetingsHandler)
	pb.RegisterAgeServiceServer(grpcServer, ageHandler)
	pb.RegisterSlowServiceServer(grpcServer, slowHandler)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
