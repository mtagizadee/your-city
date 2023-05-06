package users

import (
	"your-city/packages/db"
	"your-city/packages/users/dto"
	"your-city/packages/users/models"
)

type usersService struct {}

func (service *usersService) Create(d *dto.CreateUserDto) (*models.User, error) {
  db := db.GetDB()
  // add create logic here

  res := db.Create(&models.User{
		Name: d.Name,
		Surname: d.Surname,
		Email: d.Email,
		Password: d.Password,
	})

  if res.Error != nil {
    return nil, res.Error // add normal error validation, at least check if the email already exists or not
  }

	var created models.User
	db.Last(&created)

  return &created, nil
}

func (service *usersService) GetById(id int) (*models.User, error) {
  var user models.User
  
	db := db.GetDB()
	res := db.First(&user, id)
  
  if res.Error != nil {
    return nil, res.Error // add normal error validation, check if it is exactly not found error
  }

  return &user, nil
}

