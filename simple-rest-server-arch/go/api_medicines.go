package goapipkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"example/user/gorestapi/models"
)

type MedicinesApiController struct {
	service MedicinesApiServicer
}

func NewMedicinesApiController(s MedicinesApiServicer) Router {
	controller := &MedicinesApiController{
		service: s,
	}
	return controller
}

func (c *MedicinesApiController) Routes() Routes {
	return Routes{
		{
			Name:        "GetMedicines",
			Method:      strings.ToUpper("Get"),
			Pattern:     "/api/v1/medicines",
			HandlerFunc: c.GetMedicines,
		},
		{
			Name:        "AddMedicine",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/api/v1/medicines",
			HandlerFunc: c.AddMedicine,
		},
		{
			Name:        "SingleMedicine",
			Method:      strings.ToUpper("Get"),
			Pattern:     "/api/v1/medicines/{name}",
			HandlerFunc: c.GetMedicineByName,
		},
	}
}

func (c *MedicinesApiController) GetMedicines(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetMedicines(r.Context())
	if err != nil {
		fmt.Println("Controller:GetMedicines error")
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (c *MedicinesApiController) AddMedicine(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var medicine models.Medicine
	json.Unmarshal(reqBody, &medicine)

	result, err := c.service.AddMedicine(r.Context(), medicine)
	if err != nil {
		fmt.Println("Controller:AddMedicine error")
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (c *MedicinesApiController) GetMedicineByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	result, err := c.service.GetMedicineByName(r.Context(), name)
	if err != nil {
		fmt.Println("Controller:GetMedicineByName error")
		return
	}
	json.NewEncoder(w).Encode(result)
}
