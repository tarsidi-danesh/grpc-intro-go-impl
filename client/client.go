package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tarsidi-danesh/grpc-intro-go/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	con, err := grpc.Dial("localhost:60707", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer con.Close()

	client := proto.NewGreeterClient(con)

	req := &proto.HelloRequest{
		Name:    "Tarsidi",
		Address: "Ceger",
	}

	res, err := client.SayHello(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling SayHello RPC: %v", err)
	}

	log.Printf("Response from server: %v", res.Message)
}
