package model

import "gorm.io/gorm"


type Invetation struct{
	gorm.Model
	InvetationID uint 
	Message   string `gorm:"notNull;size:20"`
	EventID uint `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID uint 
	Accepted bool `gorm:"notNull"`
}