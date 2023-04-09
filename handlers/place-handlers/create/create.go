package createHandler

import (
	createPlace "github.com/Cyber1112/dijkstra-go/controllers/place-controllers/create"
	util "github.com/Cyber1112/dijkstra-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service createPlace.Service
}

func NewHandlerCreateStudent(service createPlace.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateStudentHandler(ctx *gin.Context) {

	var input createPlace.InputCreatePlace
	ctx.ShouldBindJSON(&input)

	_, errCreateStudent := h.service.CreatePlaceService(&input)

	switch errCreateStudent {

	case "CREATE_PLACE_CONFLICT_409":
		util.APIResponse(ctx, "Place name already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_PLACE_FAILED_403":
		util.APIResponse(ctx, "Create new place account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Created new place successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
