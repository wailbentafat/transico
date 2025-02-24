package Types

import "time"


type JwtType struct {
	Id   string `json:"id"`
	Date  time.Time `json:"date"`
	
}
