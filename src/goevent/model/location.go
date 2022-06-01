package model


type Location struct {
	LocationID  uint   	
	Name        string  `gorm:"notNull;size:20"`
	Postcode    string  `gorm:"notNull;size:7"`
	City        string  `gorm:"notNull;size:20"`
	Adress      string  `gorm:"notNull;size:20"`
	MaxCapacity uint    `gorm:"notNull;check: max_capacity>=1"`
	Latitude    float64 `gorm:"notNull"`
	Longitude   float64 `gorm:"notNull"`

}
