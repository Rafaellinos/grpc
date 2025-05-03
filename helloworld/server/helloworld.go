package main

import (
	"context"
	"log"
	pb "github.com/Rafaellinos/grpc/helloworld/proto"
	"strconv"
)

// implementa/sobre escreve o SayHello
func (s *Server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Greet call received")
	return &pb.HelloReply{
		Message: "Hello, " + request.Name,
	}, nil
}

func (s *Server) Greet(ctx context.Context, request *pb.GreetMessage) (*pb.GreetReply, error) {
	log.Printf("Greet received!!!")
	return &pb.GreetReply{
		Message: "Your name is " + request.Person.Name + " and your age is " + strconv.Itoa(int(request.Person.Age)) + " and your email is " + request.Person.Email,
	}, nil
}
