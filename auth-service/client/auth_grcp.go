package client

import (
	"context"
	"fmt"
	"log"

	"github.com/qhmd/gitforgits/shared/models"
	authpb "github.com/qhmd/gitforgits/shared/proto/auth-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	Client authpb.AuthServiceClient
	conn   *grpc.ClientConn
}

func NewAuthServiceClient(url string) *AuthServiceClient {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to auth-service: %v", err)
	}
	return &AuthServiceClient{
		Client: authpb.NewAuthServiceClient(conn),
		conn:   conn,
	}
}

func (c *AuthServiceClient) SendAuth(ctx context.Context, user *models.Auth) error {
	fmt.Println("isi user : ", user)
	result, err := c.Client.CreateAuth(ctx, &authpb.CreateAuthRequest{
		Id:       int32(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	})
	if err != nil {
		return err
	}
	fmt.Println("hasil users : ", result)

	return nil
}
