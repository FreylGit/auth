package main

import (
	"context"
	"fmt"
	desc "github.com/FreylGit/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen:%d", grpcPort)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server")
	}
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*emptypb.Empty, error) {
	log.Printf("request:%+v", req)
	return &emptypb.Empty{}, nil
}
