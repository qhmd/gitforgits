package client

import (
	"fmt"

	"github.com/qhmd/gitforgits/shared/models"
	usersproto "github.com/qhmd/gitforgits/shared/proto/users-proto"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UsersServiceClient struct {
	Client usersproto.UsersServiceClient
	conn   *grpc.ClientConn
}

func NewUsersServiceClient(url string) *UsersServiceClient {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &UsersServiceClient{
		Client: usersproto.NewUsersServiceClient(conn),
		conn:   conn,
	}
}

func (c *UsersServiceClient) UpdateUsers(ctx context.Context, user *models.Auth, id int32) {
	fmt.Printf("isi dari user update %v, dan id nya %v\n", user, id)
	result, err := c.Client.UpdateUser(ctx, &usersproto.UserRequestData{
		Id:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	})
	if err != nil {
		fmt.Println("isi error nya", err)
	}
	fmt.Println("hasil users di results users : ", result)
}

func (c *UsersServiceClient) DeleteUser(ctx context.Context, id int32) error {
	result, err := c.Client.DeleteUser(ctx, &usersproto.DeleteUserRequest{Id: id})
	if err != nil {
		fmt.Println("isi error nya", err)
		return err
	}
	fmt.Println("hasil users di results users : ", result)
	return nil
}
