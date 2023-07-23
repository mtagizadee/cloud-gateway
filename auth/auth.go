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
	
	res, err := _http.Post(base + "/login", body, map[string]string{});
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
	}, map[string]string{})
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

type VerifyDto struct {
	AccessToken string `binding:"required"`
}

func Verify(c *gin.Context) {
	body := c.MustGet("body").(VerifyDto)

	fmt.Println(body.AccessToken)
	res, err := _http.Post(base + "/verify", body, map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", body.AccessToken),
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

func Ping(c *gin.Context) {
	res, err := _http.Get(base + "/ping", map[string]string{})
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