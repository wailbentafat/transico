package middelwear

import (
	"databaseservice/pkj/jwt"
	"log"
	"strings"
	

	"github.com/gin-gonic/gin"
)



type AdminMiddleware struct {
      
}
func NewAdminmiddelweare()*AdminMiddleware{
	return &AdminMiddleware{}
}

func (r *AdminMiddleware) AdminMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		bearer:= c.GetHeader("Authorization")
		log.Println(bearer)
		if bearer == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		bearerSlice := strings.Split(bearer, "Bearer ")
		if len(bearerSlice) != 2 {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		token := bearerSlice[1]

	
		_, err := jwt.Getjwt(token)
		//use the token to get id and check if the admin
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()


		

	}

}