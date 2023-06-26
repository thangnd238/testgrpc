package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"testgrpc/proto"
)

const (
	port = 50001
)

type greeterServer struct {
	proto.UnimplementedGreeterServer
}

func (g greeterServer) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	name := in.GetName()
	return &proto.HelloReply{
		Message: "Hello" + name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &greeterServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
