package service

import (
	"errors"
	"fmt"
	"strings"
)

type ValidationError struct {
	Fields []string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("There was an error validating the following attributes: %s", strings.Join(e.Fields, ", "))
}

var ValError error = errors.New("Validation Error, Some attributes may not be right")
var NotFoundError error = errors.New("Vehicle with given ID does not exist")
