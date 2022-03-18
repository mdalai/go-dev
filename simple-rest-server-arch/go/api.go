package goapipkg

import (
	"context"

	"example/user/gorestapi/models"
)

// MedicinesApiServicer defines the api actions for the MedicinesApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type MedicinesApiServicer interface {
	AddMedicine(context.Context, models.Medicine) (ImplResponse, error)
	GetMedicines(context.Context) (ImplResponse, error)
	GetMedicineByName(context.Context, string) (ImplResponse, error)
}
