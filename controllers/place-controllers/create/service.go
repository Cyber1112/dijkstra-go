package createPlace

import "github.com/Cyber1112/dijkstra-go/models"

type Service interface {
	CreatePlaceService(input *InputCreatePlace) (*models.Place, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreatePlaceService(input *InputCreatePlace) (*models.Place, string) {
	place := models.Place{
		Name:      input.Name,
		Longitude: input.Longitude,
		Latitude:  input.Latitude,
	}

	resultCreatePlace, errCreatePlace := s.repository.CreatePlaceRepository(&place)

	return resultCreatePlace, errCreatePlace
}
