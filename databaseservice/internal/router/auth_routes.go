package router

import "github.com/gin-gonic/gin"

import "databaseservice/internal/handler"



func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)
	auth.POST("/verify", handler.Verify)
	auth.Use(middelware.JwtMiddleware())
}

