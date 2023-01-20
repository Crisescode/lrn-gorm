package model

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func DBClient() (*gorm.DB, error) {
	dsn := "user:password@tcp(ip:port)/database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
