package handlers

import (
	"fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate logic
	sum := 0
	for i := 0; i < 1000; i++ {
		sum += i
	}
	fmt.Fprintf(w, "Sum: %d", sum)
}
