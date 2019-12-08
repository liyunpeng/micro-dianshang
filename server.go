package main
import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"

	//"github.com/micro/go-grpc"
	"shopping/user/handler"
	"shopping/user/model"
	user "shopping/user/proto/user"
	"shopping/user/repository"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {

	db, err := CreateConnection()
	defer db.Close()

	db.AutoMigrate(&model.User{})

	if err != nil {
		log.Fatalf("connection error : %v \n", err)
	}

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",
		}
	})

	/*
		数据仓库对外提供了数据的增删改成
		数据仓库需要一个数据库连接db, 通过此连接来操作数据库，实现增删改查
	 */
	repo := &repository.User{db}

	service := micro.NewService(
	//service := grpc.NewService(
		micro.Name("user5"),
		micro.Registry(reg),
		//micro.Version("latest"),
	)

	service.Init()

	/*
		向服务注入包含了各种业务方法的业务结构体
		handler.User{repo} 构造出了这个业务的结构体
		这个注册的中甲的业务名UserService对应了url请求地主里的路径
		http://localhost:8080/user/UserService/Register
	 */
	user.RegisterUserServiceHandler(service.Server(), &handler.User{repo})
	//user.RegisterUserServiceHandler(service.Server(), new(handler.User{repo}))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.user", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.user", service.Server(), subscriber.Handler)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
