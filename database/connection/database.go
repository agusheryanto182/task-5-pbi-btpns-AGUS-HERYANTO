package connection

import (
	"fmt"
	"log"

	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(c *models.Global) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name)

	fmt.Printf("Connecting to database with DSN: %s\n", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Connected to the database")

	return db
}
