package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hi from Gin on Vercel!"})
	})
	router.ServeHTTP(w, r)
}
