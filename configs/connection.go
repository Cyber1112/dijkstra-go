package configs

import (
	"github.com/Cyber1112/dijkstra-go/models"
	util "github.com/Cyber1112/dijkstra-go/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	databaseURI := util.GodotEnv("DATABASE_URI")

	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	logrus.Info("Connection to Database Successfully")

	if err != nil {
		logrus.Fatal(err.Error())
	}

	err = db.AutoMigrate(&models.Place{})

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
