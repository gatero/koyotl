package main

import (
	"app/profile"
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
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProfileServiceServer(s, &profile.RPC{})
	reflection.Register(s)

	e := s.Serve(lis)
	log.Fatalf("failed to serve: %v", e)
	if e != nil {
		log.Fatalf("failed to serve: %v", e)
	}

	log.Printf("listen:%v", port)
}
