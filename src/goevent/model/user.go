package model

import "gorm.io/gorm"


type User struct{
	gorm.Model
	UserID   uint 
	Name 	string `gorm:"notNull;size:20"`
	Email 	string `gorm:"notNull;size:20"`
	Passwort 	string `gorm:"notNull;size:20"`
	Invetations []Invetation `gorm:"foreignKey:InvetationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}