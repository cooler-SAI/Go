package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log request details
		fmt.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

func main() {
	fmt.Println("web_Gin running...")

	r1 := gin.Default()
	r1.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin, it's Server 1")
	})

	r1.GET("/about", func(c *gin.Context) {
		c.String(http.StatusOK, "About us page Server 1")
	})

	r1.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "login Server 1")
	})

	r2 := gin.Default()
	r2.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin, it's Server 2")
	})

	r2.GET("/about", func(c *gin.Context) {
		c.String(http.StatusOK, "About us page Server 2")
	})

	r2.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "login Server 2")
	})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := r1.Run(":8080"); err != nil {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := r2.Run(":8081"); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}
