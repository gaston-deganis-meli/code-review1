package repository

import "errors"

var ExistingVehicleError error = errors.New("A vehicle with same ID already exists")
var NotFoundError error = errors.New("There was no vehicle with matching attributes")
var NotFound error = errors.New("Could not find vehicle with ID in DB")
var BrandNotFound error = errors.New("There are no vehicles with matching brand")
