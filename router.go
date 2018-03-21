package main

import (
	"github.com/gorilla/mux"
	"gitlab.com/john.ellis1392/todos/server"
)

// NewRouter - Create Main Router for App
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.Handle("/todos", server.TodoRouter())

	router.Handle("/api", apiRouter)
	return router
}

// Example Route Specification Structure

// func NewRouter() *mux.Router {
// 	router := mux.NewRouter().StrictSlash(true)
// 	for _, route := range routes {
// 		var handler http.Handler
//
// 		// Wrap handler function in custom logger middleware
// 		handler = route.HandlerFunc
// 		handler = Logger(handler, route.Name)
//
// 		router.Methods(route.Method).
// 			Path(route.Pattern).
// 			Name(route.Name).
// 			Handler(handler)
// 	}
//
// 	return router
// }

// package main
//
// import (
// 	"net/http"
// )
//
// type Route struct {
// 	Name        string
// 	Method      string
// 	Pattern     string
// 	HandlerFunc http.HandlerFunc
// }
//
// type Routes []Route
//
// var routes = Routes{
// 	Route{
// 		"Index",
// 		"GET",
// 		"/",
// 		Index,
// 	},
//
// 	Route{
// 		"TodoIndex",
// 		"GET",
// 		"/todos",
// 		TodoIndex,
// 	},
//
// 	Route{
// 		"TodoShow",
// 		"GET",
// 		"/todos/{todoId}",
// 		TodoShow,
// 	},
//
// 	Route{
// 		"TodoCreate",
// 		"POST",
// 		"/todos",
// 		TodoCreate,
// 	},
// }
