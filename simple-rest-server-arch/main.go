package main

import (
	"log"
	"net/http"

	goapipkg "example/user/gorestapi/go"
	//handler "example/user/gorestapi/handlers"
)

func main() {

	log.Printf("Server started")

	medService := goapipkg.NewMedicinesApiService()

	medController := goapipkg.NewMedicinesApiController(medService)

	r := goapipkg.NewRouter(medController)

	// r := mux.NewRouter().StrictSlash(true)
	// r.HandleFunc("/allmedicines", handler.AllMedicines) // ----> To request all medicines
	// r.HandleFunc("/medicines/{name}", handler.SingleMedicine)                   // ----> To request a specific medicines
	// r.HandleFunc("/medicines", handler.MedicinesToBuy).Methods("POST")          // ----> To add  new medicine
	// r.HandleFunc("/medicines/{name}", handler.UpdateMedicine).Methods("PUT")    // ----> To update a medicine
	// r.HandleFunc("/medicines/{name}", handler.DeleteMedicine).Methods("DELETE") // ----> Delete a medicine

	log.Fatal(http.ListenAndServe(":8118", r))

}
