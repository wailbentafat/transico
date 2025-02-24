package handler

import (
	"databaseservice/internal/service"
	Types "databaseservice/pkj/types"
	
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AuthHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	

}
type AuthHandlerImpl struct {

	authservice *service.Authserviceimpl
}
func Newauthhandler(authservice *service.Authserviceimpl,	)(	*AuthHandlerImpl){
	return &AuthHandlerImpl{
		
		authservice: authservice,
	}
}


func (h *AuthHandlerImpl)login (c *gin.Context)  {
	var user Types.User
	if err:=c.ShouldBindJSON(&user );err!=nil{
    fmt.Println("err in json format")
	 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	 return
	}
	if err:=h.authservice.Login	(&user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"heyy"})
	}
	//token 
	 token :=1000000

	c.JSON(http.StatusAccepted,gin.H{"token ": token })




}
func (h *AuthHandlerImpl)register(c *gin.Context)  {
	var user Types.User
	if err:=c.ShouldBindJSON(&user );err!=nil{
    fmt.Println("err in json format")
	 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	 return
	}
	if err:=h.authservice.Register	(&user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"hey"})
	}
	//token 
	 token :=1000000

	c.JSON(http.StatusAccepted,gin.H{"token ": token })




}
func (h *AuthHandlerImpl)Virify(c *gin.Context)  {
	var user Types.User
	if err:=c.ShouldBindJSON(&user );err!=nil{
    fmt.Println("err in json format")
	 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	 return
	}
	if err:=h.authservice.Virify(&user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	//token 
	 token :=1000000

	c.JSON(http.StatusAccepted,gin.H{"token ": token })




}



