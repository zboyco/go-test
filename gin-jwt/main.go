package main

import (
	"go-test/gin-jwt/middleware/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "this is index",
		})
	})
	r.POST("/Login", func(ctx *gin.Context) {
		token, err := jwt.GenerateToken(123)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
	data := r.Group("/data")
	data.Use(jwt.Auth())
	{
		data.GET("/ok", func(ctx *gin.Context) {
			userID := ctx.MustGet("UserID").(int)
			ctx.JSON(http.StatusOK, gin.H{
				"UserID": userID,
			})
		})
	}
	r.Run()
}
