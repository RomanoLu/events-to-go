package service

import (
	"github.com/RomanoLu/events-to-go/src/goevent/db"
	"github.com/RomanoLu/events-to-go/src/goevent/model"
	log "github.com/sirupsen/logrus"
)

func CreateInventation(invetation *model.Invetation) error {
	result := db.DB.Create(invetation)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new invetation with ID %v in database.", invetation.ID)
	log.Tracef("Stored: %v", invetation)
	return nil
}

func GetInvetationByUserID(userid uint) ([]model.Invetation, error) {
	var invetations []model.Invetation

	m := make(map[string]interface{})
	m["user_id"] = userid
	result:= db.DB.Where(m).Find(&invetations)
	// db.DB.Find(&invetations).Where("userid = ?", userid)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", invetations)
	return invetations, nil

}
func GetInvetationById(id uint) (*model.Invetation, error) {
	var invetation *model.Invetation
	result := db.DB.Find(&invetation, id)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", invetation)
	return invetation, nil
}

func AcceptInvetation(invetationid uint) (*model.Invetation, *model.Event, error) {
	
	//Update Invetation and make accepted true
	existingInvetation, err := GetInvetationById(invetationid)
	if existingInvetation == nil || err != nil {
		log.Error("This User do not have a Invetation")
		return existingInvetation,nil,  err
	}
	existingInvetation.Accepted = true
	result := db.DB.Save(existingInvetation)
	if result.Error != nil {
		return nil, nil, result.Error
	}
	entry := log.WithField("ID", invetationid)
	entry.Info("Successfully updated invetation.")
	entry.Tracef("Updated: %v", existingInvetation)
	
	//Add Participant to Event
	user, err := GetUserById(existingInvetation.UserID)
	if err != nil {
		log.Error("This User do not have a Invetation")
		return nil, nil, result.Error
	}
	event, err := AddParticipants(existingInvetation.EventID, user)
	if err != nil {
		log.Error("Error trying to add Participant")
		return nil, nil, result.Error
	}
	return existingInvetation,event, nil
	

}
