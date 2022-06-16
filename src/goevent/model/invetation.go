package model




type Invetation struct{
	InvetationID uint 
	Message   string `gorm:"notNull;size:20"`
	Eventid uint
	Accepted bool `gorm:"notNull"`
}