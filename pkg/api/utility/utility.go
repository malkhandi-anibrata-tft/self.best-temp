package utility

import (
	"log"

	"github.com/go-errors/errors"
)

func CreateErrorResponse(err error) map[string]string {
	LogErrorWithStacktrace(err)
	return map[string]string{
		"Error": err.Error(),
	}
}

func LogErrorWithStacktrace(err error) {
	log.Println(err.(*errors.Error).ErrorStack())
}
