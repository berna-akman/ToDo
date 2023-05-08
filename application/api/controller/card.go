package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"to-do-api/domain/card"
	"to-do-api/presentation"
)

type CardContext struct {
	mux.Route
	http.ResponseWriter
	http.Request
}

type CardController interface {
	GetAll(req *card.GetCardRequest, ctx CardContext) (*card.Card, error)
}

type cardController struct {
	s presentation.CardService
}

func NewCardController(s presentation.CardService) CardController {
	controller := &cardController{s}

	//router.HandleFunc("/card", controller.GetAll).Methods(http.MethodGet)

	return controller
}

func (c *cardController) GetAll(req *card.GetCardRequest, ctx CardContext) (*card.Card, error) {
	boards, err := c.s.GetAll()
	if err != nil {
		return nil, err
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(ctx.ResponseWriter).Encode(boards)
	if err != nil {
		return nil, err
	}
	return &boards, nil
}
