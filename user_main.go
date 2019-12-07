package main
import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	//"github.com/micro/go-grpc"
	"shopping/user/handler"
	"shopping/user/model"
	user "shopping/user/proto/user"
	"shopping/user/repository"
)

func main() {

	db, err := CreateConnection()
	defer db.Close()

	db.AutoMigrate(&model.User{})

	if err != nil {
		log.Fatalf("connection error : %v \n", err)
	}

	repo := &repository.User{db}

	//service := gppc.NewService(
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	service.Init()

	user.RegisterUserServiceHandler(service.Server(), &handler.User{repo})

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.user", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
