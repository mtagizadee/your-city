package migration

import (
	"fmt"
	"your-city/packages/db"
	"your-city/packages/users"
)

func Migrate() {
	db.GetDB().AutoMigrate(&users.User{})
	fmt.Println("Migrated the database...")
}
