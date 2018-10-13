package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// Networking
	port = 8080
)

func main() {
	bind := fmt.Sprintf(":%v", port)
	r := mux.NewRouter()
	r.HandleFunc("/", simpleHandler)

	fmt.Printf("Starting server on %v", bind)
	err := http.ListenAndServe(bind, r)
	if err != nil {
		log.Panic(err)
	}
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
