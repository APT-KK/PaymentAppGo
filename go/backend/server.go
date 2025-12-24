package main

// to export something it should start with a capital letter
import (
	"encoding/json"
	// "fmt"
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

   if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
   }

   var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

	err := json.NewDecoder(r.Body).Decode(&req) // store request body in req variable
	if(err != nil){
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// params := &stripe.PaymentIntentParams{
		
	// }

}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := []byte("Server is healthy")
	_, err := w.Write(response) // := sets type automatically
	if err != nil {
		log.Println("Error writing response:", err)
	}
}