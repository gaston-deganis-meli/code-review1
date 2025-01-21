package service

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Fields []string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("There was an error validating the following attributes: %s", strings.Join(e.Fields, ", "))
}
