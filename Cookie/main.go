package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	r.POST("/login", Login)
	r.Run()
}

func Login(c *gin.Context) {
	var req ReqParams
	err := c.ShouldBindJSON(&req)
	// c.JSON(http.StatusOK, gin.H{
	// 	"username": req.Username,
	// 	"password": req.Password,
	// })
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if req.Username != "admin" || req.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
	}
	c.JSON(200, gin.H{
		"username": req.Username,
		"password": req.Password,
		"msg":      "login success",
	})
}
