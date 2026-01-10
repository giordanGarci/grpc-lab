package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func LoggerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)

	duration := time.Since(start)

	if err != nil {
		log.Printf("❌ METHOD: %s | DURATION: %v | ERROR: %v", info.FullMethod, duration, err)
	} else {
		log.Printf("✅ METHOD: %s | DURATION: %v", info.FullMethod, duration)
	}

	return resp, err
}

func StreamLoggerInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	start := time.Now()
	err := handler(srv, ss)
	duration := time.Since(start)

	if err != nil {
		log.Printf("❌ STREAM METHOD: %s | DURATION: %v | ERROR: %v", info.FullMethod, duration, err)
	} else {
		log.Printf("✅ STREAM METHOD: %s | DURATION: %v", info.FullMethod, duration)
	}
	return err
}
