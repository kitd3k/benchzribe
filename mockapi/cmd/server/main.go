package main

import (
	"log"
	"net/http"

	"github.com/kitd3k/benchzribe/mockapi/internal/handlers"
)

func main() {
	http.HandleFunc("/api/test", handlers.TestHandler)
	log.Println("Mock API running on :8080")
	http.ListenAndServe(":8080", nil)
}
