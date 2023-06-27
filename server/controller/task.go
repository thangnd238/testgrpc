package controller

import (
	"context"
	"testgrpc/proto"
)

type TaskController struct {
	proto.UnimplementedTaskServer
}

func NewTaskController() *TaskController {
	return &TaskController{
		UnimplementedTaskServer: proto.UnimplementedTaskServer{},
	}
}

func (g *TaskController) GetTask(context.Context, *proto.TaskRequest) (*proto.TaskReply, error) {
	return &proto.TaskReply{Message: "Day la Task"}, nil
}
