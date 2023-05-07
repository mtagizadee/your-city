package auth

import (
	"crypto/sha256"
	"fmt"
	"time"
	"your-city/packages/config"
	"your-city/packages/db"
	"your-city/packages/users"

	"github.com/golang-jwt/jwt"
)

type authService struct {}

func (service *authService) signup(dto *createUserDto) (*users.User, error) {
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

func (service *authService) login(dto *LoginUserDto) (*users.User, string, error) {
	usersService := users.GetUsersService()

	// verify the email
	user, err := usersService.GetByEmail(dto.Email)
	if err != nil { 
		return nil, "", err
	}

	hash := sha256.Sum256([]byte(dto.Password)) // hash password to verify it
	if user.Password != fmt.Sprintf("%x", hash) {
		return nil, "", fmt.Errorf("wrong password")
	}

	token, err := generateJWT(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func generateJWT(user *users.User) (string, error) {
	jwtConfig := config.GetJwtConfig()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(7 * 24 * time.Hour)
	claims["email"] = user.Email

	sToken, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return "", err
	}

	return sToken, nil
}