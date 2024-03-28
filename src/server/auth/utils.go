package auth

import (
	"fmt"
	"g5/server/db"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(bytes)
}

func GenerateToken(user_id uint, role uint) (string, error) {

	token_lifespan, err := strconv.Atoi("240")

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("PRODam$%#Ah"))

}

func LoginCheck(email string, password string, pg *gorm.DB) (string, string, int, error) {

	var err error

	u := db.User{}

	err = pg.Model(db.User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		fmt.Println("email not found in database")
		return "", "Erro", 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		fmt.Println("password does not match (ERROR in bcrypt)")
		return "", "", 0, bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	} else {
		fmt.Println("password matches")
	}

	token, err := GenerateToken(u.ID, u.Role)

	if err != nil {
		fmt.Println("error generating token")
		fmt.Println(err)
		return "", "", 0, err
	}

	return token, u.Name, int(u.ID), nil

}
