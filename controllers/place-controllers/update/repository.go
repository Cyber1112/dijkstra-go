package updatePlace

import (
	"github.com/Cyber1112/dijkstra-go/models"
	"gorm.io/gorm"
)

type Repository interface {
	UpdatePlaceRepository(input *models.Place) (*models.Place, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r repository) UpdatePlaceRepository(input *models.Place) (*models.Place, string) {
	var place models.Place
	db := r.db.Model(&place)
	errorCode := make(chan string, 1)

	place.ID = input.ID

	checkPlaceId := db.Debug().Select("*").Where("id = ?", input.ID).Find(&place)

	if checkPlaceId.RowsAffected < 1 {
		errorCode <- "UPDATE_PLACE_NOT_FOUND_404"
		return &place, <-errorCode
	}

	place.Weight = input.Weight

	updatePlace := db.Debug().Where("id = ?", input.ID).Updates(place)

	if updatePlace.Error != nil {
		errorCode <- "UPDATE_PLACE_FAILED_403"
		return &place, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &place, <-errorCode
}
