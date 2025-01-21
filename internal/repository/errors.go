package repository

import "errors"

var ExistingVehicleError error = errors.New("A vehicle with same ID already exists")
var NotFoundError error = errors.New("There was no vehicle with matching attributes")
