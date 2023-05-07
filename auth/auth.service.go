package auth

import (
	"crypto/sha256"
	"fmt"
	"your-city/packages/db"
	"your-city/packages/users"
)

type authService struct {}

func (service *authService) Signup(dto *createUserDto) (*users.User, error) {
  db := db.GetDB()
  
  hash := sha256.Sum256([]byte(dto.Password)) // hash the password
	var user users.User = users.User{
		Name: dto.Name,
		Surname: dto.Surname,
		Email: dto.Email,
		Password: fmt.Sprintf("%x", hash), // paste the hash
	}

  if err := db.Create(&user).Error; err != nil {
    return nil, err
  }

  return &user, nil
}

func (service *authService) Ligin(dto *LoginUserDto) (*users.User, error) {
	usersService := users.GetUsersService()

	// verify the email
	user, err := usersService.GetByEmail(dto.Email)
	if err != nil { 
		return nil, err
	}

	hash := sha256.Sum256([]byte(dto.Password)) // hash password to verify it
	if user.Password != fmt.Sprintf("%x", hash) {
		return nil, fmt.Errorf("wrong password")
	}

	return user, nil
}