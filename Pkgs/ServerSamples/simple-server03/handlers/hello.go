package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	response := fmt.Sprintf("Hello, %s!", name)
	fmt.Fprintln(w, response)
}
