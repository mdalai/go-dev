package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"example/user/gorestapi/models"
)

var medicines = []models.Medicine{
	{Name: "Almod Milk Tyrub", Unit: "g"},
	{Name: "Apple Salt", Unit: "mg"},
}

func AllMedicines(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllMedicines")
	json.NewEncoder(w).Encode(medicines)
}

func SingleMedicine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	for _, med := range medicines {
		if med.Name == name {
			json.NewEncoder(w).Encode(med)
		}
	}
}

func MedicinesToBuy(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var medicine models.Medicine
	json.Unmarshal(reqBody, &medicine)
	medicines = append(medicines, medicine)

	json.NewEncoder(w).Encode(medicines)

}

func DeleteMedicine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	for index, med := range medicines {
		if med.Name == name {
			medicines = append(medicines[:index], medicines[index+1:]...)
		}
	}

}

func UpdateMedicine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	for index, med := range medicines {
		if med.Name == name {
			medicines = append(medicines[:index], medicines[index+1:]...)

			var updateMedicine models.Medicine

			json.NewDecoder(r.Body).Decode(&updateMedicine)
			medicines = append(medicines, updateMedicine)
			fmt.Println("Endpoint hit: UpdateMedicine")
			json.NewEncoder(w).Encode(updateMedicine)
			return
		}
	}

}
