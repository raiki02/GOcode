package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "t3/proto"
)

func main() {
	//连接
	conn, _ := grpc.NewClient("127.0.0.1:8080")
	defer conn.Close()
	//实例化
	client := pb.NewUserInfoServiceClient(conn)
	req := &pb.UserRequest{
		Name: "raiki",
	}
	res, _ := client.GetUserInfo(context.Background(), req)
	fmt.Println(res)

}
