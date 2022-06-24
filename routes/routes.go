package routes

import (
	"api/zeus/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "Up and running..."})
	})

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", handlers.GetAllUsers)
		userRoutes.GET("/:id", handlers.GetUserByID)
		userRoutes.POST("", handlers.CreateUser)
		userRoutes.PUT("/:id", handlers.UpdateUserByID)
		userRoutes.DELETE("/:id", handlers.DeleteUserByID)
		userRoutes.PATCH("/:id", handlers.PartiallyUpdatingUser)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})
}
