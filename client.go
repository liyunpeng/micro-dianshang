package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	//proto "proto_gen"
	user "shopping/user/proto/user"
)

func main() {
	/*
		创建etcd的registry,
		从registry获取一个micro服务，并将服务初始化
		micro服务的客户端， 获取一个客制化服务，
		所有的rpc请求都是由客制化服务完成
	 */
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			// 经试验, 不用配置这些地址， 只要本地启动etcd, micro的业务提供者和使用者都可正常工作
			"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
		//micro.Name("user.client")
	)
	service.Init()
	userService := user.NewUserService("user5", service.Client())

	//注册请求rpc调用
	user1 := user.User{
		//Id:       10,  注册时用户不能给自己设置id
		Name:     "name1",
		Phone:    "18712341234",
		Password: "123456",
	}
	regiserRequest := user.RegisterRequest{
		User: &user1,
	}
	res, err := userService.Register(context.TODO(), &regiserRequest)
	if err != nil {
		fmt.Println("errro:", err)
	}else{
		fmt.Println("respons message:", res.Msg)
	}

	// 登录请求rpc调用
	loginRequest := user.LoginRequest{
		Phone:    "18712341234",
		Password: "123456",
	}
	res, err = userService.Login(context.TODO(), &loginRequest)
	if err != nil {
		fmt.Println("errro:", err)
	}else{
		fmt.Println("respons message:", res.Msg)
	}

	// 修改密码的rpc调用
	UpdatePasswordRequest := user.UpdatePasswordRequest{
		Uid:1,  //  通过id找到用户， 生产环境应该是用户名
		OldPassword:"123456",
		NewPassword:"654321",
	}
	res, err = userService.UpdatePassword(context.TODO(), &UpdatePasswordRequest)
	if err != nil {
		fmt.Println("errro:", err)
	}else{
		fmt.Println("respons message:", res.Msg)
	}

}
