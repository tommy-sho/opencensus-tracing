package service

import (
	"context"
	"fmt"

	proto "github.com/tommy-sho/opencensus-tracing/user_service/genproto"
	"golang.org/x/xerrors"
)

type UserService interface {
	SaveUser(context.Context, *proto.SaveUserRequest) (*proto.SaveUserResponse, error)
	GetUser(context.Context, *proto.GetUserRequest) (*proto.GetUserResponse, error)
	GetUsers(context.Context, *proto.GetUsersRequest) (*proto.GetUsersResponse, error)
}

type userServiceImpl struct {
	Users map[string]*proto.User
}

func NewUserService() *userServiceImpl {
	users := make(map[string]*proto.User)
	return &userServiceImpl{Users: users}
}

func (u userServiceImpl) SaveUser(ctx context.Context, in *proto.SaveUserRequest) (*proto.SaveUserResponse, error) {
	u.Users[in.User.Id] = in.User
	return &proto.SaveUserResponse{User: in.User}, nil
}

func (u userServiceImpl) GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	if v, ok := u.Users[in.Id]; ok {
		return &proto.GetUserResponse{
			User: v,
		}, nil
	}
	return nil, xerrors.New(fmt.Sprintf("user_id [%v] does not exist", in.Id))
}

func (u userServiceImpl) GetUsers(ctx context.Context, in *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	users := make([]*proto.User, len(u.Users))
	counter := 0
	for _, v := range u.Users {
		users[counter] = v
		counter++
	}
	return &proto.GetUsersResponse{
		Users: users,
	}, nil
}
