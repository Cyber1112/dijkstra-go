package createPlace

import (
	"github.com/Cyber1112/dijkstra-go/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreatePlaceRepository(input *models.Place) (*models.Place, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreatePlaceRepository(input *models.Place) (*models.Place, string) {

	var place models.Place

	db := r.db.Model(&place)
	errorCode := make(chan string, 1)

	checkPlaceExist := db.Debug().Select("*").Where("name = ?", input.Name).Find(&place)

	if checkPlaceExist.RowsAffected > 0 {
		errorCode <- "CREATE_PLACE_CONFLICT_409"
		return &place, <-errorCode
	}

	place.Name = input.Name
	place.Longitude = input.Longitude
	place.Latitude = input.Latitude

	addNewPlace := db.Debug().Create(&place)
	db.Commit()

	if addNewPlace.Error != nil {
		errorCode <- "CREATE_PLACE_FAILED_403"
		return &place, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &place, <-errorCode
}
