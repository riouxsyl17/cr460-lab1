package router

import (
	"github.com/gin-gonic/gin"
	"github.com/patricklecuyer/cr460-lab1/handlers"
)

// Init initializes the routers
func Init() *gin.Engine {

	r := gin.New()

	v1 := r.Group("/v1")
	{
		contacts := v1.Group("/contacts")
		{
			contacts.GET("/", handlers.ListContacts)
			contacts.PUT("/", handlers.CreateContact)
			contacts.POST("/:email", handlers.UpdateContact)
			contacts.GET("/:email", handlers.GetContact)
		}
	}
	return r
}
