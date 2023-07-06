package model
 
import (
	"fmt"
 
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)
 
var DB *gorm.DB
 
func ConnectDatabase() {
	dsn := "makalu.iixcp.rumahweb.net user=aprh8956_dbceria password=dbceriasakti dbname=aprh8956_dbceria port=3306 sslmode=verify-full TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		fmt.Printf("Error")
		return
	}
	db = conn
}
