package main

import (
	proto "proto_gen"
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("user.client"))
	service.Init()

	user := proto.NewUserService("user", service.Client())

	res, err := user.Hello(context.TODO(), &proto.Request{Name: "World ^_^"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Msg)
}
