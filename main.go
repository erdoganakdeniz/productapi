package main

import (
	"fmt"
	"net/http"
	"productapi/handlers"
	"time"
)
var PORT=":8080"
func main() {

	mux := http.NewServeMux()
	mux.Handle("/product", http.HandlerFunc(handlers.InsertProduct))

	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	fmt.Println("Ready to serve at", PORT)
	erd := s.ListenAndServe()
	if erd != nil {
		fmt.Println(erd)
		return
	}

}


