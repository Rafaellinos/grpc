package main

import (
	"context"
	"log"
	"strconv"
	pb "github.com/Rafaellinos/grpc/helloworld/proto"
	"strings"
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

	
	if (request.GetPeople() == nil) {
		// empty response
		return &pb.GreetReply{
		}, nil
	}

	mapa := make(map[string]string, 0) // could initialize with var m map[string]string

	var builder strings.Builder

	builder.WriteString("Wellcome ")

	first := true;

	for _, person := range request.GetPeople() {
		if (!person.IsActive) {
			continue
		}
		mapa[person.Name] = "Your name is " + person.Name + " and your age is " + strconv.Itoa(int(person.Age)) + " and your email is " + person.Email
		if (!first) {
			builder.WriteString(" And ")
		}
		builder.WriteString(person.Name)
		first = false
	}

	return &pb.GreetReply{
		People: mapa,
		WellcomeMessage: builder.String(),
	}, nil
}
