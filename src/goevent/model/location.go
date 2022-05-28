package model


type Location struct{
	Name   string `gorm:"notNull;size:20"`
	Postcode string `gorm:"notNull;size:7"`
	City string `gorm:"notNull;size:20"`
	adress string `gorm:"notNull;size:20"`
	MaxCapacity uint `gorm:"notNull;check: max_capacity>=1"`
	latitude float64  `gorm:"notNull"`
	longitude float64  `gorm:"notNull"`
}