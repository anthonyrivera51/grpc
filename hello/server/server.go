package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Josh2604/grpc-hello-world/hello"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Hello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Println("Hello Called")
	name := req.GetUser().GetName()
	age := req.GetUser().GetAge()

	result := "Hello " + name + " I'm " + age + " years old"
	res := &hello.HelloResponse{
		Message: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Server running")

	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Server Creation
	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("faile to serve: %v", err)
	}
}
