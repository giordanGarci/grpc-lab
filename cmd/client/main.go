package main

import (
	"context"
	"fmt"
	"grpc-lab/pb"
	"io"
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

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// helloClient := pb.NewHelloServiceClient(conn)

	// // Contact the server and print out its response.
	// name := "Giordan"
	// r, err := helloClient.SayHello(ctx, &pb.HelloRequest{
	// 	Name: name,
	// })
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.GetMessage())

	// ageClient := pb.NewAgeServiceClient(conn)

	// birthdate := "2003-11-17"
	// ageResp, err := ageClient.GetAge(ctx, &pb.AgeRequest{
	// 	Birthdate: birthdate,
	// })
	// if err != nil {
	// 	log.Fatalf("could not get age: %v", err)
	// }
	// log.Printf("Age: %d, Is Adult: %t", ageResp.GetAge(), ageResp.GetIsAdult())

	// ctx_slqow, cancel_slow := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel_slow()

	// slowClient := pb.NewSlowServiceClient(conn)

	// slowResp, err := slowClient.ProcessSlow(ctx_slqow, &pb.SlowRequest{})
	// if err != nil {
	// 	log.Fatalf("could not process slow operation: %v", err)
	// }
	// log.Printf("Slow Operation Result: %s", slowResp.GetResult())

	// fibClient := pb.NewFibonacciServiceClient(conn)

	// fibStream, err := fibClient.GetFibonacci(ctx, &pb.FibonacciRequest{N: 10})
	// if err != nil {
	// 	log.Fatalf("could not get fibonacci numbers: %v", err)
	// }
	// for {
	// 	fibResp, err := fibStream.Recv()
	// 	if err != nil {
	// 		break
	// 	}
	// 	log.Printf("Fibonacci Number: %d", fibResp.GetValue())
	// }

	// avgAgeClient := pb.NewAgeServiceClient(conn)
	// avgAgeStream, err := avgAgeClient.ComputeAverageAge(ctx)

	// if err != nil {
	// 	log.Fatalf("could not compute average age: %v", err)
	// }
	// ages := []int32{25, 30, 22, 28, 35}
	// for _, age := range ages {
	// 	err := avgAgeStream.Send(&pb.AverageAgeRequest{Age: age})
	// 	if err != nil {
	// 		log.Fatalf("could not send age: %v", err)
	// 	}
	// }
	// avgResp, err := avgAgeStream.CloseAndRecv()
	// if err != nil {
	// 	log.Fatalf("could not receive average age: %v", err)
	// }
	// log.Printf("Final average Age: %.1f", avgResp.GetAverage())

	chatClient := pb.NewChatServiceClient(conn)
	chatStream, err := chatClient.Chat(ctx)

	if err != nil {
		log.Fatalf("could not start chat: %v", err)
	}

	// Receive messages in a separate goroutine
	go func() {
		for {
			resp, err := chatStream.Recv()
			if err == io.EOF {
				fmt.Println("Chat ended by server")
				break
			}
			if err != nil {
				log.Fatalf("error receiving chat message: %v", err)
				break
			}
			fmt.Printf("Chat bot reply from %s: %s\n", resp.GetUser(), resp.GetReply())
		}
	}()

	// main loop to send messages
	messages := []string{"Hello!", "How are you?", "Tell me a joke.", "Goodbye!"}
	for _, msg := range messages {
		fmt.Printf("sending: %s\n", msg)
		chatStream.Send(&pb.ChatMessage{
			User: "Giordan",
			Text: msg,
		})
		time.Sleep(1 * time.Second)
	}

	chatStream.CloseSend()
	time.Sleep(500 * time.Millisecond)

}
