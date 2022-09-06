package hello

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Hello struct {
	UnimplementedHelloServicesServer
}

func (hello *Hello) SayHello(ctx context.Context, in *emptypb.Empty) (*HelloResponses, error) {
	return &HelloResponses{
		Value: "Hi",
	}, nil
}
