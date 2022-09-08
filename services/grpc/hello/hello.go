package hello

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type HelloGrpcServices struct {
	UnimplementedHelloServicesServer
}

func (hello *HelloGrpcServices) SayHello(ctx context.Context, in *emptypb.Empty) (*HelloResponses, error) {
	return &HelloResponses{
		Value: "Hi",
	}, nil
}
