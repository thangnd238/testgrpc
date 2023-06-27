package controller

import (
	"context"
	"fmt"

	"testgrpc/proto"
	"testgrpc/server/service"
)

type GreeterService interface {
	SayHello(ctx context.Context, request service.HelloRequest) (service.HelloReply, error)
}

type GreeterController struct {
	proto.UnimplementedGreeterServer
	service GreeterService
}

func NewGreeterController(service GreeterService) *GreeterController {
	return &GreeterController{
		UnimplementedGreeterServer: proto.UnimplementedGreeterServer{},
		service:                    service,
	}
}

func (g *GreeterController) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	name := in.GetName()
	fmt.Println(name)

	request := service.HelloRequest{Name: name}
	result, err := g.service.SayHello(ctx, request)
	fmt.Println(result)

	return &proto.HelloReply{Message: result.Message}, err
}
