package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	var err = http.ListenAndServe("localhost:8080", nil)
	if(err != nil){
		log.Fatal("Error starting server:", err)
	}
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {

}