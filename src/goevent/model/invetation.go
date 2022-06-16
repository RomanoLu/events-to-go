package model




type Invetation struct{
	InvetationID uint 
	Message   string `gorm:"notNull;size:20"`
	Accepted bool `gorm:"notNull"`
}