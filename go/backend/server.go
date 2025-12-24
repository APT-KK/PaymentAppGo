package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/", handleHealth)
	var err error = http.ListenAndServe("localhost:8080", nil)
	if(err != nil){
		log.Fatal("Error starting server:", err)
	}
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
 fmt.Println("hiiii")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := []byte("Server is healthy")
	_, err := w.Write(response) // := sets type automatically
	if err != nil {
		log.Println("Error writing response:", err)
	}
}