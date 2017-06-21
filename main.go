package main

import (
  "fmt"
  "net/http"
  "os"
  //"github.com/johnellis1392/todos/server"
  //"github.com/johnellis1392/todos/util"
)


type RouteHandler struct {
  Path string
  State map[interface{}]interface{}
}


type Config struct {
  Port string
  Addr string
}


func getConfig() Config {
  var port = os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }

  var addr = os.Getenv("ADDR")
  if addr == "" {
    addr = "127.0.0.1"
  }

  return Config {
    Port: port,
    Addr: addr,
  }
}


func getHandler(path string, initialState map[interface{}]interface{}) RouteHandler {
  state := make(map[interface{}]interface{})
  for key, value := range initialState {
    state[key] = value
  }

  return RouteHandler{
    Path: path,
    State: state,
  }
}


func addressString(address, port string) string {
  return fmt.Sprintf("%s:%s", address, port)
}


func (handler RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  //fmt.Println(handler.State["key"])
  fmt.Fprintf(w, "Hello, World!\n")
}


func main() {
  config := getConfig()
  fmt.Printf("Starting server on port %s...\n", config.Port)

  //state["key"] = "Soemthing"
  //delete(state["key"])
  handler := getHandler("/", nil)

  address := addressString(config.Addr, config.Port)
  server := http.Server{
    Addr:    address,
    Handler: handler,
  }

  server.ListenAndServe()
}
