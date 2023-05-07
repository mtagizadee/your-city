package users

import (
	"your-city/packages/db"
)

type User struct {
  Id int `gorm:"primaryKey;autoIncrement" json:"id"`
  Name string `gorm:"not null" json:"name"`
  Surname string `gorm:"not null" json:"surname"`
  Email string `gorm:"unique;not null" json:"email"`
  Password string `gorm:"not null" json:"-"` // password should not be returned in the response
}

type usersService struct {}

func (service *usersService) GetById(id int) (*User, error) {
  var user User
  
	db := db.GetDB()
  if err := db.First(&user, id).Error; err != nil {
    return nil, err
  }

  return &user, nil
}

