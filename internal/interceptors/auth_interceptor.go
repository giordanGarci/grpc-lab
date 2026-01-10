package interceptors

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("Metadata is not provided")
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	tokens := md.Get("authorization")
	if len(tokens) == 0 || tokens[0] != "your-auth-token" {
		fmt.Println("Invalid auth token:", tokens)
		return nil, status.Errorf(codes.PermissionDenied, "invalid auth token")
	}
	fmt.Println("Auth token validated successfully")
	return handler(ctx, req)

}
