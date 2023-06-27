package service

import "context"

type Service struct{}

func NewGreeterService() Service {
	return Service{}
}

type HelloRequest struct {
	Name string
}

type HelloReply struct {
	Message string
}

func (s Service) SayHello(ctx context.Context, request HelloRequest) (HelloReply, error) {
	return HelloReply{Message: "Xin Chao " + request.Name}, nil
}
