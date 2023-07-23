package auth

import (
	"net/http"
	"packages/gateway/_http"

	"github.com/gin-gonic/gin"
)

var base = "http://localhost:3001"

type LoginDto struct {
	Email string `binding:"required,email"`
	Password string `binding:"required,min=8,max=32"`
}

type authLoginDto struct {
	Email string 
	Password string
	AppId int
	CompanyId int
}

func Login(c *gin.Context) {
	body := c.MustGet("body").(LoginDto)
	
	res, err := _http.Post(base + "/login", authLoginDto{
		Email: body.Email,
		Password: body.Password,
	});
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resBody, err := _http.Read(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusForbidden, gin.H{"error": resBody["error"]})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": resBody["accessToken"],
	})	
}

func Signup(c *gin.Context) {}

func Verify(c *gin.Context) {}

func Ping(c *gin.Context) {}