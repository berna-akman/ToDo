package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"to-do-api/pg/domain/board"
	"to-do-api/pg/presentation"
)

type BoardController interface {
	GetAllBoards(e echo.Context) error
	GetBoardByID(e echo.Context) error
	CreateBoard(e echo.Context) error
	UpdateBoard(e echo.Context) error
	DeleteBoard(e echo.Context) error
}

type boardController struct {
	s presentation.BoardService
}

func NewBoardController(s presentation.BoardService, e *echo.Echo) BoardController {
	controller := &boardController{s}

	e.GET("/pg/board", controller.GetAllBoards)
	e.GET("/pg/board/:id", controller.GetBoardByID)
	e.POST("/pg/board", controller.CreateBoard)
	e.PUT("/pg/board/:id", controller.UpdateBoard)
	e.DELETE("/pg/board/:id", controller.DeleteBoard)

	return controller
}

func (c *boardController) GetAllBoards(e echo.Context) error {
	boards, err := c.s.GetAllBoards()
	if err != nil {
		panic(err)
	}
	e.Response().Header().Set("Content-Type", "application/json")
	return e.JSON(http.StatusOK, boards)
}

func (c *boardController) GetBoardByID(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	b, err := c.s.GetBoardByID(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	e.Response().Header().Set("Content-Type", "application/json")
	return e.JSON(http.StatusOK, b)
}

func (c *boardController) CreateBoard(e echo.Context) error {
	b := board.Board{}
	err := e.Bind(&b)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.s.CreateBoard(b)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, b)
}

func (c *boardController) UpdateBoard(e echo.Context) error {
	b := &board.Board{}
	if err := e.Bind(b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.s.UpdateBoard(*b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, b)
}

func (c *boardController) DeleteBoard(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.s.DeleteBoard(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, nil)
}
