package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "github.com/tarsidi-danesh/grpc-intro-go/grpc/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	log.Println("receiving request: ", req.Name, req.Address)
	responseMessage := "Hello " + req.Name + " from " + req.Address

	response := proto.HelloReply{
		Message: responseMessage,
	}

	return &response, nil
}

func main() {
	fmt.Println("GRpc server is running")

	list, err := net.Listen("tcp", "0.0.0.0:60606")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
