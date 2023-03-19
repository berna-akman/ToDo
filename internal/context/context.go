package context

import (
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

type Context struct {
	mux.Route
	http.Request
	http.ResponseWriter
}

type Application struct {
	Server mux.Router
}

type Controller struct {
	application *Application
	routes      *mux.Router
}

func (a *Application) NewController(instance interface{}, paths ...string) *Controller {
	controller := reflect.ValueOf(instance)
	if controller.Kind() != reflect.Ptr {
		panic("Controller instance is not a pointer")
	}
	if len(paths) == 0 {
		paths = []string{""}
	}
	route := mux.NewRouter()
	return &Controller{
		application: a,
		routes:      route,
	}
}

//
//type Route struct {
//	controller *Controller
//	path       string
//	httpMethod string
//	methodName string
//	method     reflect.Value
//	models     []interface{}
//}
//
//func (c *Controller) Route(path string) *Route {
//	return &Route{
//		controller: c,
//		path:       path,
//	}
//}

//
//func (c *Controller) Route(w http.ResponseWriter, r *http.Request) error {
//	return &Route{
//		controller: c,
//		path:       path,
//	}
//}

//type Router struct {
//	NotFoundHandler         Handler
//	MethodNotAllowedHandler Handler
//	routes                  []*mux.Route
//	namedRoutes             map[string]*mux.Route
//	KeepContext             bool
//}
//
//type Route struct {
//	handler     AppHandler
//	buildOnly   bool
//	name        string
//	err         error
//	namedRoutes map[string]*Route
//}
