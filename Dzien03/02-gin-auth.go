package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func MiddlewareExample() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Query("token")
		if token != "ABCD" {
			c.AbortWithStatus(401)
			return
		}
		c.Writer.Header().Set("App-ID", "KursALX")
		c.Writer.Header().Set("X-Request-ID", uuid.NewV4().String())
	}
}

func main() {

	router := gin.Default()
	router.Use(MiddlewareExample())

	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "qwerty",
	}))
	authorized.GET("/secret", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message1": "OK"})
	})

	router.GET("/private", gin.BasicAuth(gin.Accounts{
		"admin": "qwerty123",
	}), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	router.Run()

}
