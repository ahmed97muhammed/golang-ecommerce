package routes

import (
	"golang_auth/handlers"
	"github.com/gin-gonic/gin"
	"golang_auth/middleware" 
)

func RegisterRoutes(r *gin.Engine) {
	// مسارات التسجيل والدخول
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// مسار جلب المنتجات مع التحقق من المصادقة
	r.GET("/products", middleware.AuthMiddleware(), handlers.GetProducts)
}

