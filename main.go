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

var userMap map[string]string

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
	userRepositoryCB := persistence.NewUserRepository(cb)
	userServiceCB := presentation.NewUserService(userRepositoryCB)
	userMap, err = getUserIDNameMap(userServiceCB)
	if err != nil {
		panic(err)
	}
	// TODO: investigate why userMap not applied after defining boardServiceCB
	boardServiceCB := presentation.NewBoardService(boardRepositoryCB, userMap)
	controller.NewBoardController(boardServiceCB, e)
	controller.NewUserController(userServiceCB, e)

	err = e.Start(":8080")
	if err != nil {
		panic(err)
	}
}

func getUserIDNameMap(s presentation.UserService) (map[string]string, error) {
	users, err := s.GetAllUsers()
	if err != nil {
		return nil, err
	}

	userMap = make(map[string]string)
	for _, v := range *users {
		userMap[v.ID] = v.Name
	}

	return userMap, nil
}
