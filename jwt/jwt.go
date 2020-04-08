package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"hackathon/model"
	"time"
)


var jwtKey =[]byte("a_secret_key")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	exprationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := Claims{
		UserId:         user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exprationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "NcuHomer",
			Subject: "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func ParseToken(tokenString string) (*jwt.Token, *Claims, error){
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (i interface{}, err error) {
			return jwtKey, nil
		})

	return token, claims, err
}