package model

import (
	"time"
	"gorm.io/gorm"
)

type Type string
const (
TRANSFERRED Type = "OPEN"
IN_PROCESS Type = "INVETATION_NEEDED"
)
type Event struct{
	gorm.Model
	EventID uint `gorm:"primaryKey"`
	Titel     string `gorm:"notNull;size:30"`
	Description string `gorm:"notNull;size:80"`
	Begin   time.Time `gorm:"notNull;size:20"`
	End   time.Time `gorm:"notNull;size:20"`
	MaxNumberOfParticipants   uint `gorm:"notNull; check:max_number_of_participants>=1"`
	participants   []User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Location   Location `gorm:"notNull;size:20"`
	Type   Type `gorm:"notNull;type:ENUM('OPEN','INVETATION_NEEDED')"`
	Host   User `gorm:"embedded;embeddedPrefix:user_"`
}