package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
/*
test:
http://localhost:8080/user/UserService/Register
{"user":{"id":1, "name":"test", "phone":"110", "password":"aaaaa"}}
*/

func CreateConnection() (*gorm.DB, error) {
	//host := "192.168.0.145"
	//user := "user1"
	//dbName := "shopping"
	//password := "123456"
	//return gorm.Open("mysql", fmt.Sprintf(
	//	"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
	//	user, password, host, dbName,
	//	),
	//)
	db, err := gorm.Open(
		"mysql", "root:123456@/shopping?charset=utf8&parseTime=True&loc=Local")

	return  db, err
}