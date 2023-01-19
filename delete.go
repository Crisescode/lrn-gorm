package main

import "lr-gorm/model"

func main() {
	db, err := model.DBClient()
	if err != nil {
		panic(err)
	}
	var user model.User
	db.Where("name", "Crise").Delete(&user)
}