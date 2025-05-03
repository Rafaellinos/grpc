package main

import (
	"context"
	"log"
	pb "github.com/Rafaellinos/grpc/helloworld/proto"
)

// implementa/sobre escreve o SayHello
func (s *Server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Greet call received")
	return &pb.HelloReply{
		Message: "Hello, " + request.Name,
	}, nil
}