package main

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"lr-gorm/entity"
	"lr-gorm/model"
)


func main() {
	db, err := model.DBClient()
	if err != nil {
		panic(err)
	}

	u := new(model.User)
	result := db.Debug().Model(&model.User{}).Where("name = ?", "Crise").First(u)
	var count int64
	result.Count(&count)


	ApiUser := &entity.APIUser{}
	if err := gconv.Struct(u, ApiUser); err != nil {
		panic(err)
	}
	v, err := json.Marshal(ApiUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(v))
}