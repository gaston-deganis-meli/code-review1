package handler

import (
	"app/internal/repository"
	"app/internal/service"
	"app/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv service.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv service.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]models.VehicleDoc)
		for key, value := range v {
			data[key] = models.VehicleDoc{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vehicleToCreate := models.VehicleDoc{}
		if err := json.NewDecoder(r.Body).Decode(&vehicleToCreate); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		savedV, err := h.sv.Create(vehicleToCreate)

		if err != nil {
			if errors.Is(err, repository.ExistingVehicleError) {
				response.Error(w, http.StatusConflict, err.Error())
				return
			}

			if errors.As(err, &service.ValidationError{}) {
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}

			response.Error(w, 0, err.Error()) // Internal Error
			return
		}

		response.JSON(w, http.StatusCreated, map[string]string{
			"success": fmt.Sprintf("Vehicle with ID %d was created succesfully", savedV.Id),
		})
	}
}

func (h *VehicleDefault) FindByAttrs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := chi.URLParam(r, "color")

		yearStr := chi.URLParam(r, "year")
		year, _ := strconv.Atoi(yearStr)

		foundV, err := h.sv.FindByAttrs(color, year)
		if err != nil {
			if errors.Is(err, repository.NotFoundError) {
				response.Error(w, http.StatusNotFound, "404: Not found vehicles with matching attributes")
				return
			}
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]models.VehicleDoc)
		for key, value := range foundV {
			data[key] = models.VehicleDoc{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
