package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// دالة لتشفير كلمة المرور باستخدام bcrypt
func HashPassword(password string) (string, error) {
	// تكلفة التشفير هي معلمة من bcrypt، وتتحكم في مقدار العمل
	// تزيد التكلفة مع الزمن، وأحيانًا قد تحتاج إلى تعديلها
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(hash), nil
}

// دالة للتحقق من تطابق كلمة المرور مع التشفير المخزن
func CheckPasswordHash(password, hash string) bool {
	// مقارنة كلمة المرور المدخلة مع كلمة المرور المشفرة
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
