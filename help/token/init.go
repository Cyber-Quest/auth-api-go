package token

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func Create(email string, password string, role string, id string) string {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "asdfqwerghjkvbnio") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = email
	atClaims["role"] = role
	atClaims["password"] = password
	atClaims["id"] = id

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return ""
	}
	return string(token)
}

func Compare(email string, password string, role string, id string, tokenUser string) bool {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "asdfqwerghjkvbnio") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = email
	atClaims["role"] = role
	atClaims["password"] = password
	atClaims["id"] = id
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil || string(token) != tokenUser {
		return false
	}
	return true
}

func Extract(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func Verify(r *http.Request) (*jwt.Token, error) {
	tokenString := Extract(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
