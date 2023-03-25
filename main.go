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

func main() {

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
