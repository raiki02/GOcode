package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raiki02/simple-server03/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/hello", handlers.HelloHandler)

	port := "8080"

	fmt.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
