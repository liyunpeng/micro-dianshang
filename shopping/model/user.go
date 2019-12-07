package model
/*
	不要同步import包文件， 即不能在 本文件所在目录生成vendor目录， 柔则编译报错
 */
import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string
	Phone string `gorm:"type:char(11);`
	Password string
}
