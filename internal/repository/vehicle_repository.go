package repository

import "app/pkg/models"

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)
	Save(models.Vehicle) (models.Vehicle, error)
	FindByAttrsColorNYear(color string, year int) (v map[int]models.Vehicle, err error)
	FindByAttrsBrandNYears(brand string, from, to int) (v map[int]models.Vehicle, err error)
	AverageByBrand(brand string) (average float64, err error)
	BulkSave([]models.Vehicle) error
	UpdateSpeed(id int, newSpeed float64) error
	Delete(id int) error
	FindByDimensions(minLength, maxLength, minWidth, maxWithd float64) map[int]models.Vehicle
}
