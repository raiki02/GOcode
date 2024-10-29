package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raiki02/simple-server02/router"
)

func main() {
	r := router.InitRouter()

	port := "8080"
	fmt.Printf("Server is running on port %s \n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
