package main

import (
	proto "proto_gen"
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
)

type User struct{}

func (u *User) Hello(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("user"),
	)

	service.Init()

	proto.RegisterUserHandler(service.Server(), new(User))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
