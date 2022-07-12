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

func CreateInvetation(id uint, invetation model.Invetation) (*model.User, error) {
	existingUser, err := GetUserById(id)
	if existingUser == nil || err != nil {
		return existingUser, err
	}
	existingUser.Invetations = append(existingUser.Invetations, invetation)
	result := db.DB.Save(existingUser)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully Invetated User.")
	entry.Tracef("Invetated: %v", existingUser)
	return existingUser, nil

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
