package main

import (
	"app/profile"
	"fmt"
	"log"
	"net"

	pb "app/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	lis, e := net.Listen("tcp", port)
	if e != nil {
		log.Fatalf("failed to listen: %v", e)
	}

	s := grpc.NewServer()
	pb.RegisterProfileServiceServer(s, &profile.RPC{})
	reflection.Register(s)

	if e := s.Serve(lis); e != nil {
		log.Printf("failed to serve: %v", e)
	}

	fmt.Printf("\n\n LISTEN ON: %v\n\n", port)
}
