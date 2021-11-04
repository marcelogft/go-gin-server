package routes

import (
	"kstarter/service/controller"

	"github.com/gin-gonic/gin"
)

// Routes is the function that defines the routes for the application
func SetupRouter(handler controller.Controller) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("", handler.Get)
		v1.POST("", handler.Create)
		v1.GET("/:id", handler.GetById)
		v1.PUT("/:id", handler.Update)
		v1.DELETE("/:id", handler.Delete)
	}
	return r
}
