package models

type Order struct {
	ID     int `json:"id" gorm:"primaryKey;autoIncrement"` 
	UserID int `json:"user_id" gorm:"index;not null"`     

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;"`
}