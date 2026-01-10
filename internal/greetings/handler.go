package greetings

import (
	"context"
	"grpc-lab/pb"
)

type Handler struct {
	pb.UnimplementedHelloServiceServer
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	message := h.service.SayHello(req.GetName())
	return &pb.HelloResponse{Message: message}, nil
}
