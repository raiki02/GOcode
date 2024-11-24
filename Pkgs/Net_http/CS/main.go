package main

import (
	"fmt"
	"net/http"
)

func main() {

	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,
			`package main

import (
	"fmt"
	"net/http"
)

func main() {

	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "")
	}

	http.HandleFunc("/", f)
	http.ListenAndServe(":8080", nil)
}
`)
	}

	http.HandleFunc("/input", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Query().Get("message")
		if message == "" {
			message = "no input provided"
		}
		fmt.Fprintln(w, "input: ", message)
	})
	http.HandleFunc("/", f)
	http.ListenAndServe(":8080", nil)
}
