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

	r1.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	r1.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		c.String(http.StatusOK, "Search results for: %s", query)
	})

	r1.GET("/json", func(c *gin.Context) {
		data := map[string]string{"message": "Hello, JSON"}
		c.JSON(http.StatusOK, data)
	})

	r1.POST("/json", func(c *gin.Context) {
		var jsonData map[string]interface{}
		if err := c.BindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, jsonData)
	})

	v1 := r1.Group("/v1")
	{
		v1.GET("/posts", func(c *gin.Context) {
			c.String(http.StatusOK, "List of posts")
		})

		v1.POST("/posts", func(c *gin.Context) {
			c.String(http.StatusOK, "Create a post")
		})
	}

	r2 := gin.Default()
	r2.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin, it's Server 2")
	})

	r2.GET("/about", func(c *gin.Context) {
		LoggerMiddleware()
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
