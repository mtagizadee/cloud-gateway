package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

		certificateJSON, err := json.Marshal(body.Certificate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse the certificate"})
			c.Abort()
			return
		}

		req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/certificates/verify", bytes.NewBuffer(certificateJSON))
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create request"})
			c.Abort()
			return
		}

		client := &http.Client{}
    res, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not send request"})
			c.Abort()
			return
		}

		if res.StatusCode != http.StatusOK {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid certificate"})
			c.Abort()
			return
		}


		resBody, err := io.ReadAll(res.Body)
		fmt.Println(string(resBody))

    if err != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": "Could not read response"})
				c.Abort()
				return
    }

		// parse response
		var response map[string]interface{}
		err = json.Unmarshal(resBody, &response)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse response"})
			c.Abort()
			return
		}

		c.Set("appId", response["appId"])
		c.Set("companyId", response["companyId"])

		c.Next()
	}
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}