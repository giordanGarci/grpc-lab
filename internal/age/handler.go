package age

import (
	"context"
	"grpc-lab/pb"
)

type Handler struct {
	service *Service
	pb.UnimplementedAgeServiceServer
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAge(ctx context.Context, req *pb.AgeRequest) (*pb.AgeResponse, error) {
	age, isAdult, err := h.service.GetAge(req.GetBirthdate())
	if err != nil {
		return nil, err
	}
	return &pb.AgeResponse{Age: age, IsAdult: isAdult}, nil
}
