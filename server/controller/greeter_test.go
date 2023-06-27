package controller

import (
	"context"
	"errors"
	"testgrpc/proto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testgrpc/server/service"
)

type mockGreeterService struct {
	mock.Mock
}

func newMockGreeterService() *mockGreeterService {
	return &mockGreeterService{}
}

func (m *mockGreeterService) SayHello(ctx context.Context, request service.HelloRequest) (service.HelloReply, error) {
	args := m.Called(ctx, request)

	err := args.Error(1)
	if err != nil {
		return service.HelloReply{}, err
	}

	return args.Get(0).(service.HelloReply), nil
}

func TestGreeterController_SayHello(t *testing.T) {
	tests := []struct {
		desc     string
		service  GreeterService
		request  *proto.HelloRequest
		expected *proto.HelloReply
		err      assert.ErrorAssertionFunc
	}{
		{
			desc: "service returns error",
			service: func() GreeterService {
				s := newMockGreeterService()

				s.
					On("SayHello", mock.Anything, service.HelloRequest{Name: "Thang"}).
					Return(service.HelloReply{Message: "Xin Chao Thang"}, nil)

				return s
			}(),
			request:  &proto.HelloRequest{Name: "Thang"},
			expected: &proto.HelloReply{Message: "Xin Chao Thang"},
			err:      assert.NoError,
		},
		{
			desc: "service returns error",
			service: func() GreeterService {
				s := newMockGreeterService()

				s.
					On("SayHello", mock.Anything, service.HelloRequest{Name: "Thang"}).
					Return(service.HelloReply{}, errors.New("error"))

				return s
			}(),
			request:  &proto.HelloRequest{Name: "Thang"},
			expected: &proto.HelloReply{},
			err:      assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			c := NewGreeterController(tt.service)

			result, err := c.SayHello(context.Background(), tt.request)

			assert.Equal(t, result, tt.expected)
			tt.err(t, err)
		})
	}
}
