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

	users := user.NewUserService("user5", service.Client())

	/*
		type User struct {
			Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
			Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
			Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
			Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
			XXX_NoUnkeyedLiteral struct{} `json:"-"`
			XXX_unrecognized     []byte   `json:"-"`
			XXX_sizecache        int32    `json:"-"`
		}

	*/

	user1 := user.User{
		Id:       1,
		Name:     "name1",
		Phone:    "18712341234",
		Password: "123456",
	}

	regiserRequest := user.RegisterRequest{
		User: &user1,
	}

	res, err := users.Register(context.TODO(), &regiserRequest)

	if err != nil {
		fmt.Println("errro:", err)
	}else{
		fmt.Println("respons message:", res.Msg)
	}

}
