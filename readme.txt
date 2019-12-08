生成proto:
$cd shopping
$protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto

启动etcd
$etcd

$go run server.go database.go

$ go run client.go
respons message: 注册成功
respons message: 登录成功
respons message: name1，您的密码更新成功


测试方法：
postman 测试
post: http://localhost:8080/user/UserService/Register
body raw填写：
{"user":{name":"test", "phone":"110", "password":"aaaaa"}}

$ micro api --namespace=user5
2019-12-08 17:06:14.153868 I | [api] Registering API Default Handler at /
2019-12-08 17:06:14.156868 I | [api] HTTP API Listening on [::]:8080
2019-12-08 17:06:14.157868 I | [api] Transport [http] Listening on [::]:50538
2019-12-08 17:06:14.157868 I | [api] Broker [http] Connected to [::]:50539
2019-12-08 17:06:14.195870 I | [api] Registry [mdns] Registering node: go.micro.api-023fb098-ee63-4759-b433-edd6407ebbb4
::1 - - [08/Dec/2019:17:06:23 +0800] "POST /user/UserService/Login HTTP/1.1" 500 87 "" "PostmanRuntime/7.20.1"
