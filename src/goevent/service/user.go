package service

import (
	"github.com/RomanoLu/events-to-go/src/goevent/db"
	"github.com/RomanoLu/events-to-go/src/goevent/model"
	log "github.com/sirupsen/logrus"
)

func CreateUser(user *model.User) error {
	result := db.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new user with ID %v in database.", user.ID)
	log.Tracef("Stored: %v", user)
	return nil
}

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", users)
	return users, nil
}

func GetUserById(id uint) (*model.User, error) {
	var user *model.User
	result := db.DB.Find(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", user)
	return user, nil
}
