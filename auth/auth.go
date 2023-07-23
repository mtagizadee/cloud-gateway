package auth

import (
	"fmt"
	"net/http"
	"packages/gateway/_http"

	"github.com/gin-gonic/gin"
)

var base = "http://localhost:3001"

type LoginDto struct {
	Email string `binding:"required,email"`
	Password string `binding:"required,min=8,max=32"`
}

func Login(c *gin.Context) {
	body := c.MustGet("body").(LoginDto)
	
	res, err := _http.Post(base + "/login", body);
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

type SignupDto struct {
	Email string `binding:"required,email"`
	Password string `binding:"required,min=8,max=32"`
}

type authSignupDto struct {
	Email string 
	Password string
	AppId float64
	CompanyId float64
}

func Signup(c *gin.Context) {
	body := c.MustGet("body").(SignupDto)
	appId := c.MustGet("appId").(float64)
	companyId := c.MustGet("companyId").(float64)

	fmt.Println(appId, companyId)

	res, err := _http.Post(base + "/signup", authSignupDto{
		Email: body.Email,
		Password: body.Password,
		AppId: appId,
		CompanyId: companyId,
	})
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

	c.JSON(http.StatusOK, resBody)
}

func Verify(c *gin.Context) {}

func Ping(c *gin.Context) {}