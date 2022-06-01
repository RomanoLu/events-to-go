package model



type Invetation struct{
	InvetationID uint 
	Message   string `gorm:"notNull;size:20"`
	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Event Event `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Accepted bool `gorm:"notNull"`
}