package main

import (
	"context"
	"net"
	pb "t3/proto"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedNihaoServer
}

func (s Server) Nihao(ctx context.Context, req *pb.NihaoRequest) (*pb.NihaoResponse, error) {
	return &pb.NihaoResponse{ResponseMsg: "Hello " + req.RequestName}, nil
}

func main() {

	listen, _ := net.Listen("tcp", ":8080")
	server := grpc.NewServer()
	pb.RegisterNihaoServer(server, &Server{})

	server.Serve(listen)
}
