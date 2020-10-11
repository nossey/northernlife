package infrastructure

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Db is the global DB instance
var Db *gorm.DB

func init() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	var err error
	Db, err = gorm.Open("postgres", dataSource)

	if err != nil {
		panic(err)
	}
}
