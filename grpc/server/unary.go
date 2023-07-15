package main

import (
	"context"
	pb "go_grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParams) (*pb.ResponseMessage, error) {
	return &pb.ResponseMessage{
		Message: "Hello",
	}, nil
}
