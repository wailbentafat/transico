package jwt

import (
	Types "databaseservice/pkj/types"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)


//i have to pass teh jwt secretkey arround 

func Getjwt (token string) (*Types.JwtType, error) {
	var jwttype Types.JwtType
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
        
      
        
        fmt.Println("viriying" )
		return []byte("secret"), nil
    })
    if err != nil {
		return nil, err
    }

    
    if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
        jwttype.Id = (claims["id"].(string)) 
        jwttype.Date = claims["date"].(time.Time)
    } else {
        return nil, fmt.Errorf("invalid token")
    }

    return &jwttype,nil
}






