package model

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func DBClient() (*gorm.DB, error) {
	dsn := "pdrtwd_user:123456@tcp(10.114.31.71:13310)/ppdai_pdrtwd?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
