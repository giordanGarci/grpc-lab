package age

import (
	"context"
	"fmt"
	"grpc-lab/pb"
	"io"
)

type Handler struct {
	pb.UnimplementedAgeServiceServer
	service *Service
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

func (h *Handler) ComputeAverageAge(stream pb.AgeService_ComputeAverageAgeServer) error {
	var sum int32
	var count int32

	for {
		// 1. server try to receive a request
		req, err := stream.Recv()
		// 2. if the error is io.EOF, the client has finished sending successfully
		if err == io.EOF {
			average := float64(sum) / float64(count)
			// 3. respond and close the stream connection
			return stream.SendAndClose(&pb.AverageAgeResponse{Average: average})
		}
		if err != nil {
			return err
		}
		sum += req.GetAge()
		count++
		fmt.Printf("Received age: %d\n", req.GetAge())
	}
}
