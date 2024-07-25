package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("web_Gin running...")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})

	r.GET("/about", func(c *gin.Context) {
		c.String(http.StatusOK, "About us page")
	})

	r.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "login")
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
