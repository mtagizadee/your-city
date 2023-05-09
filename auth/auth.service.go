package auth

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"
	"your-city/packages/common"
	"your-city/packages/config"
	"your-city/packages/db"
	"your-city/packages/users"
	"your-city/packages/utils"

	"github.com/golang-jwt/jwt"
)

type authService struct {}

func (service *authService) signup(dto *createUserDto) (*users.User, *common.ErrorType) {
  db := db.GetDB()
  
  hash := sha256.Sum256([]byte(dto.Password)) // hash the password
	var user users.User = users.User{
		Name: dto.Name,
		Surname: dto.Surname,
		Email: dto.Email,
		Password: fmt.Sprintf("%x", hash), // paste the hash
	}

  if err := db.Create(&user).Error; err != nil {
    if utils.IsUniqueKeyError(err) { // for some reason gorm.ErrDuplicatedKey does not work here
			return nil, &common.ErrorType{Status: http.StatusConflict, Message: "user already exists"}
		}
		
		return nil, utils.DefaultError(err)
  }

  return &user, nil
}

func (service *authService) login(dto *LoginUserDto) (*users.User, string, *common.ErrorType) {
	usersService := users.GetUsersService()

	// verify the email
	user, err := usersService.GetByEmail(dto.Email)
	if err != nil { 
		return nil, "", err
	}

	hash := sha256.Sum256([]byte(dto.Password)) // hash password to verify it
	if user.Password != fmt.Sprintf("%x", hash) {
		return nil, "", &common.ErrorType{Status: http.StatusForbidden, Message: "wrong password"}
	}

	token, err := generateJWT(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func generateJWT(user *users.User) (string, *common.ErrorType) {
	jwtConfig := config.GetJwtConfig()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(7 * 24 * time.Hour)
	claims["email"] = user.Email

	sToken, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return "", utils.DefaultError(err)
	}

	return sToken, nil
}