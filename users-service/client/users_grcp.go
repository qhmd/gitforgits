package client

import (
	"fmt"

	authproto "github.com/qhmd/gitforgits/shared/proto/auth-proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	Client authproto.AuthServiceClient
	conn   *grpc.ClientConn
}

func NewAuthServiceClient(url string) *AuthServiceClient {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &AuthServiceClient{
		Client: authproto.NewAuthServiceClient(conn),
		conn:   conn,
	}
}
