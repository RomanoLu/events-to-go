package model

type User struct{

	UserID   uint `gorm:"primaryKey"`
	Name 	string `gorm:"notNull;size:20"`
	Email 	string `gorm:"notNull;size:20"`
	Passwort 	string `gorm:"notNull;size:20"`
	AcceptedEvents 	[]Event `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Invetations []Invetation `gorm:"foreignKey:InventationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}