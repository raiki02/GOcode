package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	proto "gomicro_test/t0/proto"
)

type greeter struct {
}

func (g *greeter) HL(ctx context.Context, req *proto.REQ, res *proto.RES) error {
	res.Name = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"))
	service.Init()
	proto.RegisterGRTHandler(service.Server(), new(greeter))
	if err := service.Run(); err != nil {
		panic(err)
	}
}
