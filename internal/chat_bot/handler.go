package chatbot

import (
	"fmt"
	"grpc-lab/pb"
	"io"
)

type Handler struct {
	pb.UnimplementedChatServiceServer
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Chat(stream pb.ChatService_ChatServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Chat ended by client")
			return nil
		}
		if err != nil {
			fmt.Println("Error receiving message:", err)
			return err
		}

		fmt.Printf("Message received from %s: %s\n", req.GetUser(), req.GetText())
		response := &pb.ChatResponse{
			User:  "Server bot",
			Reply: fmt.Sprintf("Hello %s, you said: %s", req.GetUser(), req.GetText()),
		}
		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
}
