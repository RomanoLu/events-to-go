package db

import (
	"fmt"
	"os"

	"github.com/RomanoLu/events-to-go/src/goevent/model"
	"github.com/harranali/authority"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("root:root@tcp(%s)/goevent?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_CONNECT"))
	log.Info("Using DSN for DB:", dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	log.Info("Starting automatic migrations")
	if err := DB.Debug().AutoMigrate(&model.Location{}); err != nil {
		panic(err)
	}
	if err := DB.Debug().AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}
	if err := DB.Debug().AutoMigrate(&model.Event{}); err != nil {
		panic(err)
	}
	if err := DB.Debug().AutoMigrate(&model.Invetation{}); err != nil {
		panic(err)
	}
	log.Info("Automatic migrations finished")

	auth := authority.New(authority.Options{
		TablesPrefix: "event_",
		DB:           DB,
	})
	
	//Create Roles to access the event Table
	auth.CreateRole("eventhost")
	auth.CreateRole("user")

	//Create Permissions for event table
	auth.CreatePermission("delete-event")
	auth.CreatePermission("update-event")
	auth.CreatePermission("invide-user")
	auth.CreatePermission("add-participant")
	auth.CreatePermission("get-event")

	//Add Permissions to eventhost-role
	auth.AssignPermissions("eventhost", []string{
		"delete-event",
		"update-event",
		"invide-user",
		"add-participant",
		"get-event",
	})
	//Add Permissions to user-role
	auth.AssignPermissions("user", []string{
		"get-event",
	})
	

}
