package main

import (
	"net/http"

	"packages/gateway/_http"
	"packages/gateway/auth"
	"packages/gateway/common"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	api := r.Group("/api")
	v1 := api.Group("/v1")

	_auth := v1.Group("/auth")
	_auth.POST("/login", ValidateCertificateMiddleware[auth.LoginDto](), auth.Login)
	_auth.POST("/signup", ValidateCertificateMiddleware[auth.SignupDto](), auth.Signup)
	_auth.GET("/verify", auth.Verify)

	v1.GET("/ping", ValidateCertificateMiddleware[interface{}](), ping)
	r.Run("localhost:8080")
}

func ValidateCertificateMiddleware[T comparable]() gin.HandlerFunc {
	return func (c *gin.Context) {
		var body common.BaseBodyDto[T]
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("body", body.Data)

		res, err := _http.Post("http://localhost:3000/api/v1/certificates/verify", body.Certificate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return 
		}

		if res.StatusCode != http.StatusOK {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid certificate"})
			c.Abort()
			return
		}

		resBody, err := _http.Read(res)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("appId", resBody["appId"])
		c.Set("companyId", resBody["companyId"])

		c.Next()
	}
}

func ping(c *gin.Context) {	
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}