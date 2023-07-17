package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"to-do-api/cb/domain/board"
	"to-do-api/cb/presentation"
)

type BoardController interface {
	GetAllBoards(e echo.Context) error
	GetBoardByID(e echo.Context) error
	CreateBoard(e echo.Context) error
	UpdateBoard(e echo.Context) error
	DeleteBoard(e echo.Context) error
	AddColumnToBoard(e echo.Context) error
	CreateCard(e echo.Context) error
	GetCards(e echo.Context) error
	CreateCardAssignee(e echo.Context) error
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
	e.GET("/cb/board/:id/card", controller.GetCards)
	e.POST("/cb/board/:id/card/:cardId/assignee", controller.CreateCardAssignee)

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
	request := &board.DeleteColumnRequest{
		BoardID:  e.Param("id"),
		ColumnID: e.Param("columnId"),
	}
	err := c.s.RemoveColumnFromBoard(*request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, nil)
}

func (c *boardController) CreateCard(e echo.Context) error {
	request := &board.CreateCardRequest{}
	err := e.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	boardId := e.Param("id")
	id, err := c.s.CreateCard(boardId, *request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, id)
}

func (c *boardController) GetCards(e echo.Context) error {
	request := &board.GetCardRequest{
		BoardID:    e.Param("id"),
		ColumnID:   e.QueryParam("columnId"),
		AssigneeID: e.QueryParam("assigneeId"),
	}
	cards, err := c.s.GetCards(*request)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, cards)
}

func (c *boardController) CreateCardAssignee(e echo.Context) error {
	request := &board.CreateCardAssigneeRequest{
		BoardID:    e.Param("id"),
		CardID:     e.Param("cardId"),
		AssigneeID: e.QueryParam("assigneeId"),
	}
	err := e.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.s.CreateCardAssignee(*request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, nil)
}
