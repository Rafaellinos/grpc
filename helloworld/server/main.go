package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/Rafaellinos/grpc/helloworld/proto" // pb eh alias
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func main() {
	lis, error := net.Listen("tcp", "0.0.0.0:8081") //listener

	if error != nil {
		log.Fatalf("Failed, port could be in use %v", error)
	}

	log.Printf("listening on port 8081")

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})

	if error := s.Serve(lis); error != nil {
		log.Fatalf("failed to serve: %v", error)
	}

}