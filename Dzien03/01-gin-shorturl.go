package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

/*
"og3es"
http://localhost:8080/og3es
*/
type UrlShortener struct {
	Id        string    `json:"id"`
	FullUrl   string    `json:"full_url"`
	ExpiresAt time.Time `json:"expires_at"`
}

var urlMap = make(map[string]UrlShortener)

func main() {
	router := gin.Default()

	router.POST("/api/shorten", func(c *gin.Context) {
		fullUrl := c.PostForm("url")
		if fullUrl == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Podaj URL"})
			return
		}

		shortId, err := shortid.Generate()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		var expiresAt time.Time
		expiryTime, err := time.ParseDuration(c.PostForm("expiry"))
		if err == nil && expiryTime > 0 {
			expiresAt = time.Now().Add(expiryTime)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Podaj poprawny czas wygaśnięcia"})
			return
		}

		urlMap[shortId] = UrlShortener{Id: shortId, FullUrl: fullUrl, ExpiresAt: expiresAt}

		response := UrlShortener{Id: shortId, FullUrl: "http://localhost:8080/" + shortId, ExpiresAt: expiresAt}
		c.JSON(http.StatusOK, response)

	})

	router.GET("/:id", func(c *gin.Context) {
		shortId := c.Param("id")
		url, ok := urlMap[shortId]
		if ok {
			if time.Now().After(url.ExpiresAt) {
				delete(urlMap, shortId)
				c.AbortWithStatus(http.StatusGone)
			} else {
				c.Redirect(http.StatusMovedPermanently, url.FullUrl)
			}
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})

	router.Run()
}
