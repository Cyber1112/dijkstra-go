package resultsHandler

import (
	"github.com/Cyber1112/dijkstra-go/controllers/place-controllers/results"
	util "github.com/Cyber1112/dijkstra-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service results.Service
}

func NewHandlerResultsPlace(service results.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResultsPlaceHandler(ctx *gin.Context) {

	resultsPlace, errResultsPlace := h.service.ResultsPlacesService()

	switch errResultsPlace {

	case "RESULTS_PLACES_NOT_FOUND_404":
		util.APIResponse(ctx, "Places data does not exists", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Results Places data successfully", http.StatusOK, http.MethodPost, resultsPlace)
	}
}
