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

// TodoRouter - Create new Todo Router
func TodoRouter() *mux.Router {
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

// package main
//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
//
// 	"github.com/gorilla/mux"
// )
//
// func Index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Welcome!")
// }
//
// func TodoIndex(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
//
// 	if err := json.NewEncoder(w).Encode(todos); err != nil {
// 		panic(err)
// 	}
// }
//
// func TodoShow(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	todoId := vars["todoId"]
// 	fmt.Fprintln(w, "Todo show:", todoId)
// }
//
// func TodoCreate(w http.ResponseWriter, r *http.Request) {
// 	var todo Todo
// 	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	if err := r.Body.Close(); err != nil {
// 		panic(err)
// 	}
//
// 	if err := json.Unmarshal(body, &todo); err != nil {
// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		w.WriteHeader(422) // Unprocessable Entity
// 		if err := json.NewEncoder(w).Encode(err); err != nil {
// 			panic(err)
// 		}
// 	}
//
// 	t := RepoCreateTodo(todo)
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusCreated)
// 	if err := json.NewEncoder(w).Encode(t); err != nil {
// 		panic(err)
// 	}
// }
