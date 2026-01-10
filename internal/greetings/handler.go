package greetings

import (
	"context"
	"grpc-lab/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedHelloServiceServer
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	tokens := md.Get("authorization")
	if len(tokens) == 0 || tokens[0] != "your-auth-token" {
		return nil, status.Errorf(codes.PermissionDenied, "invalid auth token")
	}

	message := h.service.SayHello(req.GetName())
	return &pb.HelloResponse{Message: message}, nil
}
