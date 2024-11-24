package grpc

import (
	"context"

	"github.com/SwanHtetAungPhyo/user_service/generated/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GRPC_INT(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil , err
	}
	return conn, nil 
}

func CheckUserExist(userId uint) (bool, error){
	conn, err := GRPC_INT("localhost:50051")
	if err != nil {
		return false, err
	}
	defer conn.Close()

	client := generated.NewUserServiceClient(conn)
	resp, err := client.GetUser(context.Background(),&generated.GetUserRequest{
		Id: uint32(userId),
	})
	if err != nil {
		return false, err
	}
	
	return resp.Exist, nil
}