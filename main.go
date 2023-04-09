package main

import (
	"github.com/Cyber1112/dijkstra-go/configs"
	route "github.com/Cyber1112/dijkstra-go/routes"
	util "github.com/Cyber1112/dijkstra-go/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := SetupRouter()

	log.Fatal(router.Run(":" + util.GodotEnv("GO_PORT")))
}

func SetupRouter() *gin.Engine {

	db := configs.Connection()

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	route.InitPlaceRoutes(db, router)

	return router
}
