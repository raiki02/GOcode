package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	response := fmt.Sprintf("Hello , %s!", name)
	fmt.Fprintln(w, response)
}
