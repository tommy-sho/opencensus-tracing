package service

import (
	"context"
	"reflect"
	"sort"
	"testing"

	proto "github.com/tommy-sho/opencensus-tracing/user_service/genproto"
)

func Test_userServiceImpl_GetUser(t *testing.T) {
	type fields struct {
		Users map[string]*proto.User
	}
	type args struct {
		ctx context.Context
		in  *proto.GetUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.GetUserResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Users: map[string]*proto.User{"1": {
					Id:   "1",
					Name: "test太郎",
					Age:  10,
					Mail: "taro@gmail.com",
				}},
			},
			args: args{
				ctx: context.Background(),
				in: &proto.GetUserRequest{
					Id: "1",
				},
			},
			want: &proto.GetUserResponse{
				User: &proto.User{
					Id:   "1",
					Name: "test太郎",
					Age:  10,
					Mail: "taro@gmail.com",
				},
			},
			wantErr: false,
		},
		{
			name: "success",
			fields: fields{
				Users: map[string]*proto.User{},
			},
			args: args{
				ctx: context.Background(),
				in: &proto.GetUserRequest{
					Id: "1",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userServiceImpl{
				Users: tt.fields.Users,
			}
			got, err := u.GetUser(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userServiceImpl_GetUsers(t *testing.T) {
	type fields struct {
		Users map[string]*proto.User
	}
	type args struct {
		ctx context.Context
		in  *proto.GetUsersRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.GetUsersResponse
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				Users: map[string]*proto.User{
					"1": {
						Id:   "1",
						Name: "test太郎",
						Age:  10,
						Mail: "taro@gmail.com",
					},
					"2": {
						Id:   "2",
						Name: "test二郎",
						Age:  20,
						Mail: "jiro@gmail.com",
					}},
			},
			args: args{
				ctx: nil,
				in:  &proto.GetUsersRequest{},
			},
			want: &proto.GetUsersResponse{
				Users: []*proto.User{
					{
						Id:   "1",
						Name: "test太郎",
						Age:  10,
						Mail: "taro@gmail.com",
					},
					{
						Id:   "2",
						Name: "test二郎",
						Age:  20,
						Mail: "jiro@gmail.com",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userServiceImpl{
				Users: tt.fields.Users,
			}
			got, err := u.GetUsers(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				got.Users = mySort(got.Users)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func mySort(f []*proto.User) []*proto.User {
	sort.Slice(f, func(i, j int) bool { return f[i].Id < f[j].Id })
	return f
}
