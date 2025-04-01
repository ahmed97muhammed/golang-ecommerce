package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"os"
)

// هيكل البيانات للتوثيق
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// ميدل وير للتحقق من صحة JWT في الترويسة
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// استخراج التوكن من رأس الطلب
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// إزالة كلمة "Bearer " من التوكن
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// التحقق من التوكن
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// التحقق من التوقيع باستخدام المفتاح السري من البيئة
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// إذا كان التوكن صالحًا، استمر في تنفيذ الطلب
		c.Set("username", claims.Username)
		c.Next()
	}
}
