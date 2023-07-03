package main

import (
	"github.com/labstack/echo/v4"
	"to-do-api/cb/application/api/controller"
	"to-do-api/cb/infrastructure/persistence"
	"to-do-api/cb/presentation"
	_ "to-do-api/pg/application/api/controller"
	_ "to-do-api/pg/infrastructure/config"
	_ "to-do-api/pg/infrastructure/persistence"
	_ "to-do-api/pg/presentation"
)

func main() {

	/*cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	pg, err := persistence2.Connect(*cfg)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	boardRepository := persistence2.NewBoardRepository(pg)
	boardService := presentation2.NewBoardService(boardRepository)
	controller2.NewBoardController(boardService, e)*/

	e := echo.New()
	cb, err := persistence.ConnectCB()
	boardRepositoryCB := persistence.NewBoardRepository(cb)
	boardServiceCB := presentation.NewBoardService(boardRepositoryCB)
	controller.NewBoardController(boardServiceCB, e)

	err = e.Start(":8080")
	if err != nil {
		panic(err)
	}
}
