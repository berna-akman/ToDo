package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"to-do-api/cb/domain/board"
	"to-do-api/cb/presentation"
	"to-do-api/internal/errors"
)

type BoardController interface {
	GetAllBoards(e echo.Context) error
	GetBoardByID(e echo.Context) error
	CreateBoard(e echo.Context) error
	UpdateBoard(e echo.Context) error
	DeleteBoard(e echo.Context) error
	AddColumnToBoard(e echo.Context) error
	CreateCard(e echo.Context) error
}

type boardController struct {
	s presentation.BoardService
}

func NewBoardController(s presentation.BoardService, e *echo.Echo) BoardController {
	controller := &boardController{s}

	e.GET("/cb/board", controller.GetAllBoards)
	e.GET("/cb/board/:id", controller.GetBoardByID)
	e.POST("/cb/board", controller.CreateBoard)
	e.PUT("/cb/board/:id", controller.UpdateBoard)
	e.DELETE("/cb/board/:id", controller.DeleteBoard)

	e.POST("/cb/board/:id/column", controller.AddColumnToBoard)
	e.DELETE("/cb/board/:id/column/:columnId", controller.RemoveColumnFromBoard)

	e.POST("/cb/board/:id/card", controller.CreateCard)

	return controller
}

func (c *boardController) GetAllBoards(e echo.Context) error {
	boards, err := c.s.GetAllBoards()
	if err != nil {
		return err
	}
	e.Response().Header().Set("Content-Type", "application/json")
	return e.JSON(http.StatusOK, boards)
}

func (c *boardController) GetBoardByID(e echo.Context) error {
	var b *board.Board
	id := e.Param("id")

	b, err := c.s.GetBoardByID(id)
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

	id, err := c.s.CreateBoard(b)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, id)
}

func (c *boardController) UpdateBoard(e echo.Context) error {
	b := &board.Board{}
	if err := e.Bind(b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := e.Param("id")
	if err := c.s.UpdateBoard(id, *b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, nil)
}

func (c *boardController) DeleteBoard(e echo.Context) error {
	id := e.Param("id")
	err := c.s.DeleteBoard(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, nil)
}

func (c *boardController) AddColumnToBoard(e echo.Context) error {
	column := board.Column{}
	err := e.Bind(&column)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	boardId := e.Param("id")
	id, err := c.s.AddColumnToBoard(boardId, column)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, id)
}

func (c *boardController) RemoveColumnFromBoard(e echo.Context) error {
	boardId := e.Param("id")
	columnId := e.Param("columnId")
	err := c.s.RemoveColumnFromBoard(boardId, columnId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, nil)
}

func (c *boardController) CreateCard(e echo.Context) error {
	card := board.Card{}
	err := e.Bind(&card)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	boardId := e.Param("id")
	id, err := c.s.CreateCard(boardId, card)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, id)
}
