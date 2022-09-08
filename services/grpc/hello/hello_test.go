package hello

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	RegisterHelloServicesServer(server, &HelloGrpcServices{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestSayHello(t *testing.T) {

	ctx := context.Background()

	conn, err := grpc.DialContext(
		ctx, "",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(dialer()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewHelloServicesClient(conn)
	res, err := client.SayHello(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	require.Equal(t, "Hi", res.Value, "Should return Hi")
}
