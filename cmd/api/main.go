package main

import (
	"ecommerce/config"
	"ecommerce/internal/database"
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
	//db connect
	db,err:=database.DbConnect(config.GetDbConfig())
	if err != nil {
		log.Fatal("Error connecting db", err)
		
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! I am invisible")
	})
	// Start server
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
