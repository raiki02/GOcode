package router

import (
	"github.com/gorilla/mux"
	"github.com/raiki02/simple-server02/handlers"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/hello/{name}", handlers.HelloHandler).Methods("GET")
	return r
}
