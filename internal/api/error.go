package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Status int
	Errors any
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v", e.Errors)
}

func HttpErrorHandler(err error, c echo.Context) {
	NewResponse(c).SetErrors(err).Send()
}

type ValidationError map[string][]string

func (ve ValidationError) Error() string {
	return fmt.Sprintf("%v", map[string][]string(ve))
}

func (ve ValidationError) IsFail() bool {
	return len(ve) > 0
}

func (ve ValidationError) Add(field string, messages ...string) {
	value, exists := ve[field]
	if !exists {
		value = []string{}
	}

	ve[field] = append(value, messages...)
}

func (ve ValidationError) AddError(field string, err error) {
	if err != nil {
		ve.Add(field, err.Error())
	}
}
