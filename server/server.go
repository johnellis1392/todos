package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"gitlab.com/john.ellis1392/todos/model"

	"github.com/gorilla/mux"
)

// ListTodos - Fetch a list of todos
func ListTodos(w http.ResponseWriter, r *http.Request) {
	// ...
}

// FetchTodo - Get a specific todo
func FetchTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var todo model.Todo
	const readerLimit int64 = 1048576
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, readerLimit))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	fmt.Println("Received Todo:", todo)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

// NewRouter - Create new Todo Router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").
		Path("/api/todos").
		Name("ListTodos").
		HandlerFunc(ListTodos)

	router.Methods("GET").
		Path("/api/todos/{todoId}").
		Name("FetchTodo").
		HandlerFunc(FetchTodo)

	return router
}
