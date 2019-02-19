package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prongbang/go-testcov/pm25/di"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	h := http.NewServeMux()

	// Route
	di.ProvidePmRoute().Initial(h)

	log.Printf("Server starting port 8080")

	http.ListenAndServe(":8080", h)
}
