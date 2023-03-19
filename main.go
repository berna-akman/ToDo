package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"to-do-api/application/api/controller"
	"to-do-api/infrastructure/config"
	"to-do-api/infrastructure/persistence"
	"to-do-api/presentation"
)

// @title To Do Api Swagger
// @version 1.0
// @description To Do
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http https

//type App struct {
//	Router *mux.Router
//}

func main() {

	// TODO: testlere odaklan, daha fazla test veya validasyon koyabilir miyiz? controller nasıl test edebilirim vs, dockerize ve deploy edelim - docker compose

	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	pg, err := persistence.Connect(*cfg)
	if err != nil {
		panic(err)
	}

	boardRepository := persistence.NewBoardRepository(pg)
	cardRepository := persistence.NewCardRepository(pg)
	boardService := presentation.NewBoardService(boardRepository)
	cardService := presentation.NewCardService(cardRepository)
	router := mux.NewRouter()
	controller.NewBoardController(boardService, router)
	controller.NewCardController(cardService, router)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
