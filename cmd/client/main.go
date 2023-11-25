package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kentro-lab/cloudy-neigh/proto/api"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCloudyNeighClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Index(ctx, &pb.IndexRequest{})
	if err != nil {
		log.Fatalf("Index error: %v", err)
	}
	log.Printf("Response: %s", r)
}
