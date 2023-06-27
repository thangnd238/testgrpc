package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"testgrpc/proto"
	"testgrpc/server/controller"
	"testgrpc/server/service"
)

const (
	port = 50001
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	greeterService := service.NewGreeterService()
	greeterController := controller.NewGreeterController(greeterService)

	taskController := controller.NewTaskController()

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, greeterController)
	proto.RegisterTaskServer(s, taskController)
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
