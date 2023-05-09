package users

import (
	"errors"
	"net/http"
	"your-city/packages/common"
	"your-city/packages/db"
	"your-city/packages/utils"

	"gorm.io/gorm"
)

type User struct {
  Id int `gorm:"primaryKey;autoIncrement" json:"id"`
  Name string `gorm:"not null" json:"name"`
  Surname string `gorm:"not null" json:"surname"`
  Email string `gorm:"unique;not null" json:"email"`
  Password string `gorm:"not null" json:"-"` // password should not be returned in the response
}

type usersService struct {}

func (service *usersService) GetById(id int) (*User, *common.ErrorType) {
  var user User
  
	db := db.GetDB()
  if err := db.First(&user, id).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return nil, &common.ErrorType{Status: http.StatusNotFound, Message: "user with this id is not found"}
    }

    return nil, utils.DefaultError(err)
  }

  return &user, nil
}

func (service *usersService) GetByEmail(email string) (*User, *common.ErrorType) {
  var user User
  
  db := db.GetDB()
  if err := db.Where("email = ?", email).First(&user).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return nil, &common.ErrorType{Status: http.StatusNotFound, Message: "user with this email is not found"}
    }
    
    return nil, utils.DefaultError(err)
  }

  return &user, nil

}

func GetUsersService() *usersService {
  return service
}