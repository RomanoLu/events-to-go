package model

import "gorm.io/gorm"



type Invetation struct{
	gorm.Model
	InvetationID uint 
	Message   string `gorm:"notNull;size:20"`
	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Event Event `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Accepted bool `gorm:"notNull"`
}