package updatePlace

import "github.com/Cyber1112/dijkstra-go/models"

type Service interface {
	UpdatePlaceService(input *InputUpdatePlace) (*models.Place, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdatePlaceService(input *InputUpdatePlace) (*models.Place, string) {

	place := models.Place{
		ID:        input.ID,
		Name:      input.Name,
		Longitude: input.Longitude,
		Latitude:  input.Latitude,
	}

	resultUpdatePlace, errUpdatePlace := s.repository.UpdatePlaceRepository(&place)

	return resultUpdatePlace, errUpdatePlace
}
