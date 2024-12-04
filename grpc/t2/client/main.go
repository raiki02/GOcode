package main

import (
	"context"
	"fmt"
	pb "t3/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	client := pb.NewNihaoClient(conn)

	res, _ := client.Nihao(context.Background(), &pb.NihaoRequest{RequestName: "aaasdf"})
	fmt.Println(res.GetResponseMsg())
}
