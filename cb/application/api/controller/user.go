package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"to-do-api/cb/domain/user"
	"to-do-api/cb/presentation"
)

type UserController interface {
	GetAllUsers(e echo.Context) error
	CreateUser(e echo.Context) error
	GetUserByID(e echo.Context) error
}

type userController struct {
	s presentation.UserService
}

func NewUserController(s presentation.UserService, e *echo.Echo) UserController {
	controller := &userController{s}

	e.GET("/cb/user", controller.GetAllUsers)
	e.POST("/cb/user", controller.CreateUser)
	e.GET("/cb/user/:id", controller.GetUserByID)

	return controller
}

func (c *userController) GetAllUsers(e echo.Context) error {
	users, err := c.s.GetAllUsers()
	if err != nil {
		return err
	}
	e.Response().Header().Set("Content-Type", "application/json")
	return e.JSON(http.StatusOK, users)
}

func (c *userController) CreateUser(e echo.Context) error {
	user := user.User{}
	err := e.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := c.s.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, id)
}

func (c *userController) GetUserByID(e echo.Context) error {
	var u *user.User
	id := e.Param("id")

	u, err := c.s.GetUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	e.Response().Header().Set("Content-Type", "application/json")
	return e.JSON(http.StatusOK, u)
}
