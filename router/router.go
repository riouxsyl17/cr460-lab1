package router

import "github.com/gin-gonic/gin"

// Init initializes the routers
func Init() *gin.Engine {
	r := gin.New()
	r.Static("/", "./static")

	return r
}
