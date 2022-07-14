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
