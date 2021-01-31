package main

import (
	"context"
	"log"
	"net"

	pb "github.com/trif0lium/stellari-sh/apps/danbooru/generated/apis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedGalleryServer
}

func (s *server) Posts(ctx context.Context, in *pb.PostsRequest) (*pb.PostsResponse, error) {
	if len(in.Tags) == 0 {
		return nil, status.Error(codes.FailedPrecondition, "Tags should not be empty.")
	}

	return &pb.PostsResponse{
		Posts: nil,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGalleryServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
