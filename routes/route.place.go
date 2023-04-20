package route

import (
	createPlace "github.com/Cyber1112/dijkstra-go/controllers/place-controllers/create"
	"github.com/Cyber1112/dijkstra-go/controllers/place-controllers/results"
	updatePlace "github.com/Cyber1112/dijkstra-go/controllers/place-controllers/update"
	createHandler "github.com/Cyber1112/dijkstra-go/handlers/place-handlers/create"
	resultsHandler "github.com/Cyber1112/dijkstra-go/handlers/place-handlers/results"
	updateHandler "github.com/Cyber1112/dijkstra-go/handlers/place-handlers/update"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPlaceRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Place
	*/
	createPlaceRepository := createPlace.NewRepositoryCreate(db)
	createPlaceService := createPlace.NewServiceCreate(createPlaceRepository)
	createPlaceHandler := createHandler.NewHandlerCreatePlace(createPlaceService)

	resultsPlaceRepository := results.NewRepositoryResults(db)
	resultsPlaceService := results.NewServiceResults(resultsPlaceRepository)
	resultsPlaceHandler := resultsHandler.NewHandlerResultsPlace(resultsPlaceService)

	updatePlaceRepository := updatePlace.NewRepositoryUpdate(db)
	updatePlaceService := updatePlace.NewServiceUpdate(updatePlaceRepository)
	updatePlaceHandler := updateHandler.NewHandlerUpdatePlace(updatePlaceService)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/places", createPlaceHandler.CreatePlaceHandler)
	groupRoute.GET("/places", resultsPlaceHandler.ResultsPlaceHandler)
	groupRoute.PUT("/places/:id", updatePlaceHandler.UpdatePlaceHandler)
}
