package service

import "app/pkg/models"

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)
	Create(vDoc models.VehicleDoc) (savedV models.Vehicle, err error)
	FindByAttrsColorNYear(color string, year int) (v map[int]models.Vehicle, err error)
	FindByAttrsBrandNYears(brand string, from, to int) (map[int]models.Vehicle, error)
	AverageByBrand(brand string) (average float64, err error)
	BulkCreate([]models.VehicleDoc) error
	UpdateSpeed(id int, newSpeed float64) error
	Delete(id int) error
	FindByDimensions(minLength, maxLength, minWidth, maxWithd float64) (map[int]models.Vehicle, error)
	FindByWeight(minW, maxW float64) (map[int]models.Vehicle, error)
}
