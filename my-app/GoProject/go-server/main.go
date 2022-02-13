package main

// imported packages
import (
	"fmt"
	"log"
	"net/http"

	"go-server/router"
)

// main function

func main() {

	r := router.Router()

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", r))
}
