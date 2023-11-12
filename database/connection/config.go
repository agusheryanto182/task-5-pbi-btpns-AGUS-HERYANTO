package connection

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:Root1234!@tcp(localhost:3306)/go_personalize_user?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
