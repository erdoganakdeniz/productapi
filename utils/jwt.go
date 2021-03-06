package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

func CreateJWTToken(data map[string]interface{}) (string, error) {
	mapClaims:=make(jwt.MapClaims,len(data))
	for key, val := range data {
		mapClaims[key]=val
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,mapClaims)
	tokenSecret:=os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(tokenSecret))
}
