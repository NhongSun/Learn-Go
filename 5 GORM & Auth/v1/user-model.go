package main

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

func createUser(db *gorm.DB, user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func login(db *gorm.DB, user *User) (string, error) {
	selectedUser := new(User)
	// get user from email
	result := db.Where("email = ?", user.Email).First(selectedUser)

	if result.Error != nil {
		return "", result.Error
	}

	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))

	if err != nil {
		return "", err
	}

	// return jwt token if passed
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = selectedUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
