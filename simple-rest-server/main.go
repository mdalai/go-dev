package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	handler "example/user/gorestapi/handlers"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/allmedicines", handler.AllMedicines)                         // ----> To request all groceries
	r.HandleFunc("/medicines/{name}", handler.SingleMedicine)                   // ----> To request a specific grocery
	r.HandleFunc("/medicines", handler.MedicinesToBuy).Methods("POST")          // ----> To add  new grocery to buy
	r.HandleFunc("/medicines/{name}", handler.UpdateMedicine).Methods("PUT")    // ----> To update a grocery
	r.HandleFunc("/medicines/{name}", handler.DeleteMedicine).Methods("DELETE") // ----> Delete a grocery

	log.Fatal(http.ListenAndServe(":8118", r))

}
