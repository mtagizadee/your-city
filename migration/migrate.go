package migration

import (
	"fmt"
	"your-city/packages/countries"
	"your-city/packages/db"
	"your-city/packages/users"
)

func Migrate() {
	db := db.GetDB()
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&countries.Country{})
	
	fmt.Println("Migrated the database...")
}
