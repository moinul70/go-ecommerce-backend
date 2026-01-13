package main

import (
	"ecommerce/config"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	cnf := config.GetEnv()
	portStr := cnf.HttpPort
	port := strconv.Itoa(portStr)
	fmt.Printf("Starting server on :%s\n", port)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! I am invisible")
	})
	// Start server
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
