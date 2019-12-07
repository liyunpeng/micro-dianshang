生成proto:
$cd shopping
$protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto

$consul agent -dev

$go run main.go database.go

$micro api --namespace=go.micro.srv

测试方法：
postman 测试
post: http://localhost:8080/user/UserService/Register
body raw填写：
{"user":{"id":1, "name":"test", "phone":"110", "password":"aaaaa"}}