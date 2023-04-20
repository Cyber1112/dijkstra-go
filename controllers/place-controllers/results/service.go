package results

import "github.com/Cyber1112/dijkstra-go/models"

type Service interface {
	ResultsPlacesService() (*[]models.Place, string)
}

type service struct {
	repository Repository
}

func NewServiceResults(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResultsPlacesService() (*[]models.Place, string) {

	results, errResults := s.repository.ResultsPlacesRepository()

	return results, errResults
}

func (s *service) DijkstraService(nodeA string) (*[]models.Place, string) {

	results, errResults := s.repository.ResultsPlacesRepository()

	return results, errResults
}
