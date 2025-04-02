package handlers

import (
	"golang_auth/models"
	"golang_auth/utils"
	"golang_auth/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte(utils.GetEnv("JWT_SECRET")) // تحميل المفتاح من المتغيرات البيئية

func Register(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, please provide valid JSON"})
		return
	}

	// Validation: Check if username and password are provided
	if newUser.Username == "" || newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	// Validation: Check if username is already taken
	existingUser, err := models.GetUserByUsername(newUser.Username)
	if err == nil && existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is already taken"})
		return
	}

	// Validation: Check if username is already taken
	existingUser2, err := models.GetUserByEmail(newUser.Email)
	if err == nil && existingUser2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already taken"})
		return
	}

	// Hashing the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}
	newUser.Password = string(hashedPassword)

	// Insert the user into the database
	if err := newUser.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while registering user"})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// التحقق من اسم المستخدم وكلمة المرور
	user, err := models.GetUserByUsername(creds.Username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// إنشاء JWT
	claims := &middleware.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})
}
