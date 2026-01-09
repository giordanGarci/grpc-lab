package main

import (
	"context"
	"fmt"
	"grpc-lab/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	initTime := time.Now()
	defer func() {
		fmt.Printf("Execution time: %s\n", time.Since(initTime))
	}()

	fmt.Println("starting client grpc")

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	helloClient := pb.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	name := "Giordan"
	r, err := helloClient.SayHello(ctx, &pb.HelloRequest{
		Name: name,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	ageClient := pb.NewAgeServiceClient(conn)

	birthdate := "2003-11-17"
	ageResp, err := ageClient.GetAge(ctx, &pb.AgeRequest{
		Birthdate: birthdate,
	})
	if err != nil {
		log.Fatalf("could not get age: %v", err)
	}
	log.Printf("Age: %d, Is Adult: %t", ageResp.GetAge(), ageResp.GetIsAdult())

	ctx_slqow, cancel_slow := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel_slow()

	slowClient := pb.NewSlowServiceClient(conn)

	slowResp, err := slowClient.ProcessSlow(ctx_slqow, &pb.SlowRequest{})
	if err != nil {
		log.Fatalf("could not process slow operation: %v", err)
	}
	log.Printf("Slow Operation Result: %s", slowResp.GetResult())

}
