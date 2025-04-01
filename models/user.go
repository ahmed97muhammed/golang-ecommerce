package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang_auth/utils"
)

var db *sql.DB

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	// تحميل المتغيرات البيئية
	utils.LoadEnv()

	// الاتصال بقاعدة البيانات
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", 
		utils.GetEnv("DB_USER"), 
		utils.GetEnv("DB_PASSWORD"), 
		utils.GetEnv("DB_HOST"), 
		utils.GetEnv("DB_NAME"))
	
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

func (u *User) Create() error {
	_, err := db.Exec("INSERT INTO users(username, password, email) VALUES(?, ?, ?)", u.Username, u.Password, u.Email)
	return err
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, password, email FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}
