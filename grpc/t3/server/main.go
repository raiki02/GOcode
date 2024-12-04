package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "t3/proto"
)

// 操作接口
type UserInfoService struct {
	//pb.UnimplementedUserInfoServiceServer
	pb.UserInfoServiceServer
}

var (
	u = UserInfoService{}
)

func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (res *pb.UserResponse, err error) {
	name := req.Name
	if name == "raiki" {
		res = &pb.UserResponse{
			Name:  name,
			Age:   3,
			Id:    1,
			Hobby: []string{"coding", "0"},
		}
	}
	return res, nil
}

func main() {
	//监听
	addr := "127.0.0.1:8080"
	listener, _ := net.Listen("tcp", addr)
	fmt.Println("server is running on", addr)
	//实例化
	s := grpc.NewServer()
	//注册
	pb.RegisterUserInfoServiceServer(s, &u)

	s.Serve(listener)
}
