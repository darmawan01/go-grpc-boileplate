package main

import (
	"fmt"
	"go_grpc_boileplate/services/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

func init() {

}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hello.RegisterHelloServicesServer(s, &hello.HelloServices{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
