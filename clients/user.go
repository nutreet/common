package clients

import (
	"context"

	proto "github.com/nutreet/common/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	auth proto.AuthServiceClient
}

func NewUserClient(addr string) (*UserClient, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	auth := proto.NewAuthServiceClient(conn)

	return &UserClient{
		auth,
	}, nil
}

func (c *UserClient) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	resp, err := c.auth.Register(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}