package fibonacci

import (
	"grpc-lab/pb"
)

type Handler struct {
	service *Service
	pb.UnimplementedFibonacciServiceServer
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetFibonacci(req *pb.FibonacciRequest, stream pb.FibonacciService_GetFibonacciServer) error {
	ctx := stream.Context()
	n := req.GetN()
	fibNumbers := make(chan int64)

	go h.service.CalculateFibonacci(ctx, n, fibNumbers)

	for num := range fibNumbers {
		if err := stream.Send(&pb.FibonacciResponse{Value: num}); err != nil {
			return err
		}
	}

	return nil
}
