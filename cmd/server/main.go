package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/kentro-lab/cloudy-neigh/proto/api"
)

var port = flag.Int("port", 50051, "The server port")

type server struct {
	pb.UnimplementedCloudyNeighServer
}

func (*server) Index(context.Context, *pb.IndexRequest) (*pb.IndexResponse, error) {
	return &pb.IndexResponse{}, nil
}

func (*server) Search(context.Context, *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCloudyNeighServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
