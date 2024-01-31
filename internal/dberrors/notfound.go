package dberrors

import "fmt"

type NotFound struct {
	Entity string
	ID     string
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("unable to find %s with id %s", e.Entity, e.ID)
}
