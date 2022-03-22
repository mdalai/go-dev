package goapipkg

import (
	"context"
	"errors"
	"net/http"

	"example/user/gorestapi/models"
)

var medicines = []models.Medicine{
	{Name: "Almod Milk Tyrub", Unit: "g"},
	{Name: "Apple Salt", Unit: "mg"},
	{Name: "Banz", Unit: "mg"},
}

type MedicinesApiService struct {
}

func NewMedicinesApiService() MedicinesApiServicer {
	return &MedicinesApiService{}
}

func (s *MedicinesApiService) GetMedicines(ctx context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, medicines), nil
}

func (s *MedicinesApiService) AddMedicine(ctx context.Context, medicine models.Medicine) (ImplResponse, error) {
	medicines = append(medicines, medicine)
	return Response(http.StatusCreated, medicines), nil
}

func (s *MedicinesApiService) GetMedicineByName(ctx context.Context, name string) (ImplResponse, error) {
	for _, med := range medicines {
		if med.Name == name {
			return Response(200, med), nil
		}
	}
	return Response(http.StatusNotFound, nil), errors.New("Medicine not found")
}
