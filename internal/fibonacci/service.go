package fibonacci

import (
	"context"
	"time"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CalculateFibonacci(ctx context.Context, n int32, out chan int64) error {
	defer close(out)
	if n < 0 {
		return nil
	}
	var a, b int64 = 0, 1
	for i := int32(0); i < n; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case out <- a:
			a, b = b, a+b
			time.Sleep(200 * time.Millisecond)
		}
	}
	return nil
}
