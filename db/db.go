package db

import (
	"fmt"
	"your-city/packages/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func Connect() {
  databaseConfig := config.GetDatabaseConfig()
  dataBase, err := gorm.Open(mysql.Open(databaseConfig.Dsn), &gorm.Config{})
  if err != nil {
    panic("ERROR: Unable to connect to the database")
  } else { fmt.Println("Connected to the database...") }

  database = dataBase
}

func GetDB() *gorm.DB {
  return database
}
