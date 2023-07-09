package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "vinmyid_dbceria:dbceriasakti@tcp(iix101.idcloudhost.com:3306)/vinmyid_dbceria?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		fmt.Printf("Error")
		return
	}
	DB = conn
}
