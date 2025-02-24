package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)


func Createjwt (id string) string {
	// token :=jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["id"] = id
	// tokenString, err := token.SignedString([]byte("secret"))
	// if err != nil {
	// 	println(err)
	// }
	// return tokenString
	token := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		jwt.MapClaims{
			"id":   id,
			"date": time.Now().UTC().Format(time.RFC3339),
		},
	)
	
	
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		println(err)
	}
	return tokenString

}

