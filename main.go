package main

import (
	"github.com/labstack/echo/v4"
	"to-do-api/cb/application/api/controller"
	"to-do-api/cb/infrastructure/persistence"
	"to-do-api/cb/presentation"
	controller2 "to-do-api/pg/application/api/controller"
	"to-do-api/pg/infrastructure/config"
	persistence2 "to-do-api/pg/infrastructure/persistence"
	presentation2 "to-do-api/pg/presentation"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	pg, err := persistence2.Connect(*cfg)
	if err != nil {
		panic(err)
	}
	cb, err := persistence.ConnectCB()

	e := echo.New()
	boardRepository := persistence2.NewBoardRepository(pg)
	boardService := presentation2.NewBoardService(boardRepository)
	controller2.NewBoardController(boardService, e)
	//cardRepository := persistence2.NewCardRepository(pg)
	//cardService := presentation2.NewCardService(cardRepository)
	//controller2.NewCardController(cardService)

	boardRepositoryCB := persistence.NewBoardRepository(cb)
	boardServiceCB := presentation.NewBoardService(boardRepositoryCB)
	controller.NewBoardController(boardServiceCB, e)

	err = e.Start(":8080")
	if err != nil {
		panic(err)
	}
}

// TODO: board'lara default degerlerle column arrayi ekleme (column isimleri tutulacak) --- done

// TODO: column olusturma, silme

// TODO: board'ları column'a gore filtrelemek icin query'den gonder (GET)

// TODO: card board'un bir column'una ait olacak iliskinisi kuralım
