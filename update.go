package main

import "lr-gorm/model"

func main() {
	db, err := model.DBClient()
	if err != nil {
		panic(err)
	}

	db.Model(&model.User{}).Where("name", "Crise").Update("age", 28)
}