package service

import (
	"app/internal/repository"
	"app/pkg/models"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp repository.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp repository.VehicleRepository
}

func validateVehicle(v models.VehicleDoc) ([]string, error) {
	failedAttr := make([]string, 0)
	if v.Brand == "" {
		failedAttr = append(failedAttr, "Brand")
	}
	if v.Model == "" {
		failedAttr = append(failedAttr, "Model")
	}
	if v.Registration == "" {
		failedAttr = append(failedAttr, "Registration")
	}
	if v.Color == "" {
		failedAttr = append(failedAttr, "Color")
	}
	if v.FabricationYear <= 0 {
		failedAttr = append(failedAttr, "FabricationYear")
	}
	if v.Capacity <= 0 {
		failedAttr = append(failedAttr, "Capacity")
	}
	if v.MaxSpeed <= 0.0 {
		failedAttr = append(failedAttr, "MaxSpeed")
	}
	if v.FuelType == "" {
		failedAttr = append(failedAttr, "FuelType")
	}
	if v.Transmission == "" {
		failedAttr = append(failedAttr, "Transmission")
	}
	if v.Weight <= 0.0 {
		failedAttr = append(failedAttr, "Weight")
	}
	if v.Length <= 0.0 {
		failedAttr = append(failedAttr, "Length")
	}
	if v.Width <= 0.0 {
		failedAttr = append(failedAttr, "Width")
	}

	if len(failedAttr) > 0 {
		return failedAttr, ValidationError{failedAttr}
	}

	return failedAttr, nil
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) Create(vDoc models.VehicleDoc) (models.Vehicle, error) {

	// Validate
	if _, err := validateVehicle(vDoc); err != nil {
		return models.Vehicle{}, err
	}

	// Creating Model
	v := models.Vehicle{
		Id: vDoc.ID,
		VehicleAttributes: models.VehicleAttributes{
			Brand:           vDoc.Brand,
			Model:           vDoc.Model,
			Registration:    vDoc.Registration,
			Color:           vDoc.Color,
			FabricationYear: vDoc.FabricationYear,
			Capacity:        vDoc.Capacity,
			MaxSpeed:        vDoc.MaxSpeed,
			FuelType:        vDoc.FuelType,
			Transmission:    vDoc.Transmission,
			Weight:          vDoc.Weight,
			Dimensions: models.Dimensions{
				Height: vDoc.Height,
				Length: vDoc.Length,
				Width:  vDoc.Width,
			},
		},
	}

	// Save Model
	savedV, err := s.rp.Save(v)
	return savedV, err
}

func (s *VehicleDefault) FindByAttrsColorNYear(color string, year int) (v map[int]models.Vehicle, err error) {
	// Validation acá ---
	v, err = s.rp.FindByAttrsColorNYear(color, year)
	return
}

func ValidateBrandNYears(brand string, from, to int) error {
	errorFields := make([]string, 0)
	if brand == "" {
		errorFields = append(errorFields, "Brand")
	}
	if from == 0 {
		errorFields = append(errorFields, "FromYear")
	}
	if to == 0 {
		errorFields = append(errorFields, "ToYear")
	}
	if len(errorFields) > 0 {
		return ValidationError{errorFields}
	}
	return nil
}

func (s *VehicleDefault) FindByAttrsBrandNYears(brand string, from, to int) (map[int]models.Vehicle, error) {
	if err := ValidateBrandNYears(brand, from, to); err != nil {
		return map[int]models.Vehicle{}, err
	}

	foundV, err := s.rp.FindByAttrsBrandNYears(brand, from, to)
	return foundV, err
}

func (s *VehicleDefault) AverageByBrand(brand string) (average float64, err error) {
	if brand == "" {
		return average, ValidationError{[]string{"Brand"}}
	}
	average, err = s.rp.AverageByBrand(brand)

	return
}

func (s *VehicleDefault) BulkCreate(vDocs []models.VehicleDoc) error {

	vehicles := make([]models.Vehicle, 0)

	for _, vDoc := range vDocs {
		if _, err := validateVehicle(vDoc); err != nil {
			return err
		}

		v := models.Vehicle{
			Id: vDoc.ID,
			VehicleAttributes: models.VehicleAttributes{
				Brand:           vDoc.Brand,
				Model:           vDoc.Model,
				Registration:    vDoc.Registration,
				Color:           vDoc.Color,
				FabricationYear: vDoc.FabricationYear,
				Capacity:        vDoc.Capacity,
				MaxSpeed:        vDoc.MaxSpeed,
				FuelType:        vDoc.FuelType,
				Transmission:    vDoc.Transmission,
				Weight:          vDoc.Weight,
				Dimensions: models.Dimensions{
					Height: vDoc.Height,
					Length: vDoc.Length,
					Width:  vDoc.Width,
				},
			},
		}

		vehicles = append(vehicles, v)
	}

	err := s.rp.BulkSave(vehicles)
	return err
}

func (s *VehicleDefault) UpdateSpeed(id int, newSpeed float64) error {
	if id <= 0 || newSpeed <= 0 || newSpeed > 350 {
		return ValError
	}

	if err := s.rp.UpdateSpeed(id, newSpeed); err != nil {
		return NotFoundError
	}

	return nil
}

func (s *VehicleDefault) Delete(id int) error {
	if err := s.rp.Delete(id); err != nil {
		return NotFoundError
	}
	return nil
}

func (s *VehicleDefault) FindByDimensions(minLength, maxLength, minWidth, maxWithd float64) (map[int]models.Vehicle, error) {
	// if minLength == 0.0 || maxLength == 0.0 || minWidth == 0.0 || maxWithd == 0.0 {
	// 	return map[int]models.Vehicle{}, ValError
	// }

	foundV := s.rp.FindByDimensions(minLength, maxLength, minWidth, maxWithd)
	if len(foundV) == 0 {
		return map[int]models.Vehicle{}, NotFoundMatchingError
	}

	return foundV, nil
}
