package results

import (
	"github.com/Cyber1112/dijkstra-go/models"
	"gorm.io/gorm"
)

type Repository interface {
	ResultsPlacesRepository() (*[]models.Place, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryResults(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ResultsPlacesRepository() (*[]models.Place, string) {

	var places []models.Place
	db := r.db.Model(&places)
	errorCode := make(chan string, 1)

	resultsPlaces := db.Debug().Select("*").Find(&places)

	if resultsPlaces.Error != nil {
		errorCode <- "RESULTS_PLACES_NOT_FOUND_404"
		return &places, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &places, <-errorCode
}
