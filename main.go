package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prshendra/golang-rest/handlers"
)

func main() {
	r := gin.Default()
	handlers.RegisterRoutes(R)
	r.Run()
}
