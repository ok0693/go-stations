package model

import "fmt"


type ErrNotFound struct {
	NotFound string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf(e.NotFound)
}