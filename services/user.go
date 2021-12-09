package services

import (
	"context"
	"fmt"

	"github.com/thompsonmss/grpc-go/pb"
	"github.com/thompsonmss/grpc-go/pb/pb"
)

type UserServiceServer interface {
	AddUser(context.Context, *pb.User) (*pb.User, error)
	mustEmbedUnimplementedUserServiceServer()
	// AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)
}

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer)
