package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type dto struct {
	Input string
}

func RegisterRoutes(r *gin.Engine) {
	r.POST("/upper", func(c *gin.Context) {
		var data dto

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"message": `invalid input format: {"input": "your string"}`,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": strings.ToUpper(data.Input),
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/date-now", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": time.Now().String(),
		})
	})
}
