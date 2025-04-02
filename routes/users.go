package routes

import (
	"golang_auth/handlers"
	"golang_auth/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Authentication Group
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", handlers.Register)
		authGroup.POST("/login", handlers.Login)
	}

	// Product Group (requires authentication middleware)
	productGroup := r.Group("/products")
	{
		productGroup.Use(middleware.AuthMiddleware()) // Apply middleware to this group
		productGroup.GET("/", handlers.GetProducts)  // List products
	}
}
