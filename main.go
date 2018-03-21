package main

import (
	"fmt"
	"log"
	"net/http"
)

// RouteHandler - HttpHandler for nested routing
type RouteHandler struct {
	Path   string
	State  map[interface{}]interface{}
	Config Config
}

func getHandler(path string, initialState map[interface{}]interface{}) RouteHandler {
	state := make(map[interface{}]interface{})
	for key, value := range initialState {
		state[key] = value
	}

	return RouteHandler{
		Path:  path,
		State: state,
	}
}

func (handler RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func (handler RouteHandler) listenAndServe() error {
	server := http.Server{
		Addr:    handler.Config.AddressString(),
		Handler: handler,
	}

	return server.ListenAndServe()
}

func main() {
	config := EnvConfig()
	fmt.Printf("Starting server on port %s...\n", config.Port)

	handler := getHandler("/", nil)

	// address := config.AddressString()
	// server := http.Server{
	// 	Addr:    address,
	// 	Handler: handler,
	// }

	// server.ListenAndServe()

	log.Fatal(handler.listenAndServe())
}
