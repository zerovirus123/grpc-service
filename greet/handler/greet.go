package handler

import (
	"context"
	"fmt"
	greet "grpc-server/grpc-server/pb/go"
)

type Greet struct {
	greet.UnimplementedServer
}

func (d Greet) Greeting(ctx context.Context, m *greet.GreetingRequest) (*greet.GreetingResponse, error) {
	return &greet.GreetingResponse{
		Message: fmt.Sprintf("gRPC server: HI %v!", m.Name),
	}, nil
}
