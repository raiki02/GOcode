package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to the simple go server01 ! ")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	response := fmt.Sprintf("Hello , %s!", name)
	fmt.Fprintln(w, response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/hello/{name}", HelloHandler)

	port := "8080"
	fmt.Printf("server is running on port %s \n", port)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
