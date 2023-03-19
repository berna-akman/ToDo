package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
	"log"
	"net/http"
	"strconv"
	"to-do-api/domain/board"
	. "to-do-api/internal/context"
	"to-do-api/presentation"
)

type BoardController interface {
	GetAll(w http.ResponseWriter, r *http.Request) error
	GetByID(w http.ResponseWriter, r *http.Request) error
	Create(w http.ResponseWriter, r *http.Request) error
	Update(w http.ResponseWriter, r *http.Request) error
	Delete(w http.ResponseWriter, r *http.Request) error
}

type boardController struct {
	s presentation.BoardService
}

// @title Swagger To Do API
// @version 1.0
// @description This is a sample server To Do server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /

func NewBoardController(s presentation.BoardService, router *mux.Router) BoardController {
	controller := &boardController{s}

	// TODO: generic'leri kullan

	//router := Application{Server: router}
	//route := router.NewController(controller)
	router.Handle("/board", Handler(controller.GetAll)).Methods(http.MethodGet) // controller.getxxx??
	router.Handle("/board/{id:[0-9]+}", Handler(controller.GetByID)).Methods(http.MethodGet)
	router.Handle("/board", Handler(controller.Create)).Methods(http.MethodPost)
	router.Handle("/board/{id:[0-9]+}", Handler(controller.Update)).Methods(http.MethodPut)
	router.Handle("/board/{id:[0-9]+}", Handler(controller.Delete)).Methods(http.MethodDelete)

	router.HandleFunc("/swagger/*", httpSwagger.WrapHandler) // The url pointing to API definition

	return controller
}

func (c *boardController) GetAll(w http.ResponseWriter, r *http.Request) error {
	boards, err := c.s.GetAll()
	if err != nil {
		panic(err)
	}
	// TODO: ortaklanabilir aynı sey yapıldıgı icin, middleware or func ile handle edebilirsin
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(boards)
	if err != nil {
		return err
	}
	return nil
}

func (c *boardController) GetByID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	b, err := c.s.GetByID(id)
	if err != nil {
		w.WriteHeader(404)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		panic(err)
	}
	return err
}

func (c *boardController) Create(w http.ResponseWriter, r *http.Request) error {
	b := board.Board{}
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Fatal(err)
	}
	err = c.s.Create(b)
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(err)
	if err != nil {
		return err
	}
	return nil
}

func (c *boardController) Update(w http.ResponseWriter, r *http.Request) error {
	var b *board.Board
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		return err
	}

	err = c.s.Update(*b)
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		return err
	}
	return nil
}

func (c *boardController) Delete(w http.ResponseWriter, r *http.Request) error {
	var b *board.Board
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}

	err = c.s.Delete(id)
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		return err
	}
	return nil
}

//func (c *boardController) GetAll(w http.ResponseWriter, r *http.Request) {
//	boards, _ := c.s.GetAll()
//	err := Encoder(w, boards)
//	if err != nil {
//		panic(err)
//	}
//}

//func (c *boardController) GetByID(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id, _ := strconv.Atoi(vars["id"])
//	b, err := c.s.GetByID(id)
//	if err != nil {
//		w.WriteHeader(404)
//	}
//	err = Encoder(w, b)
//	if err != nil {
//		return
//	}
//	return
//}

//func (c *boardController) Create(w http.ResponseWriter, r *http.Request) {
//	b := board.Board{}
//	err := json.NewDecoder(r.Body).Decode(&b)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = c.s.Create(b)
//	err = Encoder(w, err)
//	if err != nil {
//		return
//	}
//}

//func (c *boardController) Delete(w http.ResponseWriter, r *http.Request) {
//	var b *board.Board
//	vars := mux.Vars(r)
//	id, _ := strconv.Atoi(vars["id"])
//
//	err := c.s.Delete(id)
//	if err != nil {
//		return
//	}
//	err = Encoder(w, b)
//	if err != nil {
//		return
//	}
//}
