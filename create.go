package main

import (
	"lr-gorm/model"
	"time"
)

func main() {
	db, err := model.DBClient()
	if err != nil {
		panic(err)
	}

	// 创建 User Table
	db.AutoMigrate(&model.User{})

	// 创建一条记录
	email := "zhaopanpan01@xinye.com"
	birthday := time.Date(1995, 01, 01, 10, 00,00, 0, time.Local)
	user := model.User{Name: "zhaopanpan01", Email: &email, Age: 28, BirthDay: &birthday}
	db.Create(&user)

	// 批量添加记录
	var users = []model.User{{Name: "Vicky", Age: 22}, {Name: "Bob", Age: 30}, {Name: "Tom", Age: 35}}
	db.Create(&users)
}

