package tools

import (
	"fmt"

	"lawSite/models"
)

func CreateError(errorString string, err error) *models.Error {
	var errFormat string
	if err != nil {
		errFormat = fmt.Sprintf("%s: %s", errorString, err)
	} else {
		errFormat = errorString
	}
	return &models.Error{Data: &models.ErrorData{Error: errFormat}}
}
