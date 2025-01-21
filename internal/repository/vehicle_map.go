package repository

import (
	"app/pkg/models"
	"strings"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]models.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]models.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]models.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) Save(v models.Vehicle) (models.Vehicle, error) {
	if _, ok := r.db[v.Id]; !ok {
		r.db[v.Id] = v
		return v, nil
	}
	return models.Vehicle{}, ExistingVehicleError
}

func (r *VehicleMap) FindByAttrs(color string, year int) (v map[int]models.Vehicle, err error) {
	foundV := make(map[int]models.Vehicle)
	for id, vehicle := range r.db {
		if strings.ToLower(vehicle.Color) == strings.ToLower(color) && vehicle.FabricationYear == year {
			foundV[id] = vehicle
		}
	}
	if len(foundV) > 0 {
		return foundV, nil
	}
	return foundV, NotFoundError
}
