package main

import (
	"fmt"
	"log"
	"net/http"

	"go-server/router"
)

func main() {

	r := router.Router()

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Function for handling messages
