package router

import (
	"github.com/gorilla/mux"
	"github.com/jrakhman/logger"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		//apply logger decorator
		handler := logger.NewLogger(route.HandlerFunc, route.Name)

		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)
	}

	return router
}
