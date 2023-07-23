package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginDto struct {
	Email string `binding:"required,email"`
	Password string `binding:"required,min=8,max=32"`
}

func Login(c *gin.Context) {
	body := c.MustGet("body").(LoginDto)
	// appId := c.MustGet("appId").(float64)
	// companyId := c.MustGet("companyId").(float64)

	
	fmt.Println(body)
	
}

func Signup(c *gin.Context) {}

func Verify(c *gin.Context) {}