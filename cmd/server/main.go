package main

import (
	"fmt"
	"net"

	"grpc-lab/internal/age"
	chatbot "grpc-lab/internal/chat_bot"
	"grpc-lab/internal/fibonacci"
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
		grpc.ChainUnaryInterceptor(
			interceptors.LoggerInterceptor,
			interceptors.AuthInterceptor),
		grpc.StreamInterceptor(interceptors.StreamLoggerInterceptor),
	)

	greetingsService := greetings.NewService()
	greetingsHandler := greetings.NewHandler(greetingsService)

	ageService := age.NewService()
	ageHandler := age.NewHandler(ageService)

	slowService := slow.NewService()
	slowHandler := slow.NewHandler(slowService)

	fibonacciService := fibonacci.NewService()
	fibonacciHandler := fibonacci.NewHandler(fibonacciService)

	chatHandler := chatbot.NewHandler()

	pb.RegisterHelloServiceServer(grpcServer, greetingsHandler)
	pb.RegisterAgeServiceServer(grpcServer, ageHandler)
	pb.RegisterSlowServiceServer(grpcServer, slowHandler)
	pb.RegisterFibonacciServiceServer(grpcServer, fibonacciHandler)
	pb.RegisterChatServiceServer(grpcServer, chatHandler)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
