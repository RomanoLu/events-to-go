package model


type User struct{
	UserID   uint 
	Name 	string `gorm:"notNull;size:20"`
	Email 	string `gorm:"notNull;size:20"`
	Passwort 	string `gorm:"notNull;size:20"`
}