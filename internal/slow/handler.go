package slow

import (
	"context"
	"errors"
	"grpc-lab/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedSlowServiceServer
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}
func (h *Handler) ProcessSlow(ctx context.Context, req *pb.SlowRequest) (*pb.SlowResponse, error) {
	result, err := h.service.SimulateSlowOperation(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "tempo esgotado")
		}
		return nil, status.Error(codes.Canceled, "operação abortada")
	}
	return &pb.SlowResponse{Result: result}, nil
}
