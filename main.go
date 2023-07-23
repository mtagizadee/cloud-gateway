package main

import (
	"net/http"

	"packages/gateway/_http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.Use(validateCertificate())

	api := r.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/ping", ping)
	r.Run("localhost:8080")
}

type PublicCertificate struct {
	AccessToken string `binding:"required"`
	ApplicationId int `binding:"required"`
  CertificateId string `binding:"required"`
  CreatedAt string `binding:"required"`
}

type BaseBodyDto struct {
	Certificate PublicCertificate 
}

func validateCertificate() gin.HandlerFunc {
	
	return func (c *gin.Context) {
		var body BaseBodyDto
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Missing certificate"})
			c.Abort()
			return
		}

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