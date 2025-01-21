package repository

import "errors"

var ExistingVehicleError error = errors.New("A vehicle with same ID already exists")
