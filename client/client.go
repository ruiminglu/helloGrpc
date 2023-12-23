package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"ruiming.com/helloGrpc/helloworld"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := helloworld.NewGreeterClient(conn)

	req := &helloworld.HelloRequest{
		Name: "Gopher",
	}

	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	log.Printf("Response from server: %s", resp.Greeting)
}
