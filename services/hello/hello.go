package hello

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type HelloServices struct {
	UnimplementedHelloServicesServer
}

func (hello *HelloServices) SayHello(ctx context.Context, in *emptypb.Empty) (*HelloResponses, error) {
	return &HelloResponses{
		Value: "Hi",
	}, nil
}
