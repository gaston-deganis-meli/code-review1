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
	"strings"

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

		foundV, err := h.sv.FindByAttrsColorNYear(color, year)
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

func (h *VehicleDefault) FindByAttrsBrandNYears() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get our params
		brand := chi.URLParam(r, "brand")

		yearFromStr := chi.URLParam(r, "start_year")
		yearToStr := chi.URLParam(r, "end_year")

		yearFrom, _ := strconv.Atoi(yearFromStr)
		yearTo, _ := strconv.Atoi(yearToStr)

		foundV, err := h.sv.FindByAttrsBrandNYears(brand, yearFrom, yearTo)

		// Error handling
		if err != nil {
			if errors.Is(err, repository.NotFoundError) {
				response.Error(w, http.StatusNotFound, err.Error())
				return
			}
			if errors.As(err, &service.ValidationError{}) {
				response.Error(w, http.StatusBadRequest, err.Error())
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

func (h *VehicleDefault) AverageByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		av, err := h.sv.AverageByBrand(brand)

		// Error handling
		if err != nil {
			if errors.Is(err, repository.BrandNotFound) {
				response.Error(w, http.StatusNotFound, err.Error())
				return
			}

			if errors.As(err, &service.ValidationError{}) {
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}

			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		data := fmt.Sprintf("average speed is %f", av)
		response.JSON(w, http.StatusOK, map[string]any{
			"status":  "success",
			"message": data,
		})

	}
}

func (h *VehicleDefault) BulkCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vDocs := make([]models.VehicleDoc, 0)

		// Decode body
		if err := json.NewDecoder(r.Body).Decode(&vDocs); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// Error handling
		if err := h.sv.BulkCreate(vDocs); err != nil {
			if errors.Is(err, repository.ExistingVehicleError) {
				response.Error(w, http.StatusConflict, err.Error())
				return
			}
			if errors.As(err, &service.ValidationError{}) {
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}
			response.Error(w, 0, err.Error())
			return
		}

		// Todo ok :)

		response.JSON(w, http.StatusCreated, map[string]string{
			"status":  "created",
			"message": "Successfully created vehicles",
		})
	}
}

type UpdateSpeedRequest struct {
	NewSpeed float64 `json:"new_speed"`
}

func (h *VehicleDefault) UpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, _ := strconv.Atoi(idStr)

		updateRequest := UpdateSpeedRequest{}
		if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// Error handling
		if err := h.sv.UpdateSpeed(id, updateRequest.NewSpeed); err != nil {
			if errors.Is(err, service.ValError) {
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}
			if errors.Is(err, service.NotFoundError) {
				response.Error(w, http.StatusNotFound, err.Error())
				return
			}
			// Internal Server Error
			response.Error(w, 0, err.Error())
			return
		}

		data := fmt.Sprintf("successfully updated vehicle with id %d", id)
		response.JSON(w, http.StatusOK, map[string]string{
			"success": data,
		})

	}
}

func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		if err = h.sv.Delete(id); err != nil {
			if errors.Is(err, service.NotFoundError) {
				response.Error(w, http.StatusNotFound, err.Error())
				return
			}
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		data := fmt.Sprintf("successfully deleted vehicle with id: %d", id)
		response.JSON(w, http.StatusNoContent, map[string]string{
			"success": data,
		})
	}
}

func (h *VehicleDefault) FindByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		lenRange := query.Get("length")
		widthRange := query.Get("width")

		lengthParts := strings.Split(lenRange, "-")
		widthParts := strings.Split(widthRange, "-")

		fmt.Println(lengthParts)
		fmt.Println(widthParts)

		if len(lengthParts) == 2 && len(widthParts) == 2 {
			minLength, err := strconv.ParseFloat(lengthParts[0], 64)
			if err != nil {
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}
			maxLength, err := strconv.ParseFloat(lengthParts[1], 64)
			if err != nil {
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}
			minWidth, err := strconv.ParseFloat(widthParts[0], 64)
			if err != nil {
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}
			maxWidth, err := strconv.ParseFloat(widthParts[1], 64)
			if err != nil {
				response.Error(w, http.StatusBadRequest, err.Error())
				return
			}

			foundV, err := h.sv.FindByDimensions(minLength, maxLength, minWidth, maxWidth)
			if err != nil {
				if errors.Is(err, service.NotFoundMatchingError) {
					response.Error(w, http.StatusNotFound, err.Error())
					return
				}
				if errors.Is(err, service.ValError) {
					response.Error(w, http.StatusBadRequest, err.Error())
					return
				}
				response.Error(w, 0, err.Error())
				return
			}

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
			return

		} else {
			response.Error(w, http.StatusBadRequest, "Incorrect query structure")
			return
		}
	}
}
