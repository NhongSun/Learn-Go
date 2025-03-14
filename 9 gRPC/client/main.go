package main

import (
	"context"
	"log"
	"time"

	pb "sunny-client/grpc-hello/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a new insecure credentials
	// secure hijacking of the connection
	creds := insecure.NewCredentials()

	// Create a new gRPC client connection
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := "Mike"
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// GetMessage is a method of the HelloReply struct
	log.Printf("Greeting: %s", r.GetMessage())
}
