package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// docker启动mysql: docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass123 -d 7484689f290f
func Connect() {
	d, err := gorm.Open("mysql", "root:pass123@(localhost)/bookstore?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
