package updateHandler

import (
	updatePlace "github.com/Cyber1112/dijkstra-go/controllers/place-controllers/update"
	util "github.com/Cyber1112/dijkstra-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	service updatePlace.Service
}

func NewHandlerUpdatePlace(service updatePlace.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) UpdatePlaceHandler(ctx *gin.Context) {

	var input updatePlace.InputUpdatePlace

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	input.ID = uint(id)

	ctx.ShouldBindJSON(&input)

	_, errUpdatePlace := h.service.UpdatePlaceService(&input)

	switch errUpdatePlace {

	case "UPDATE_PLACE_NOT_FOUND_404":
		util.APIResponse(ctx, "Place's data does not exist or deleted", http.StatusNotFound, http.MethodPost, nil)

	case "UPDATE_PLACE_FAILED_403":
		util.APIResponse(ctx, "Update place data failed", http.StatusForbidden, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Update place data successfully", http.StatusOK, http.MethodPost, nil)
	}
}
