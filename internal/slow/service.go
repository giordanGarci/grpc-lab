package slow

import (
	"context"
	"fmt"
	"time"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SimulateSlowOperation(ctx context.Context) (string, error) {
	done := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		done <- "Operation completed"
	}()

	select {
	case res := <-done:
		fmt.Println("Operation completed successfully")
		return res, nil
	case <-ctx.Done():
		fmt.Println("Operation cancelled")
		// Handle context cancellation
		return "", ctx.Err()
	}
}
