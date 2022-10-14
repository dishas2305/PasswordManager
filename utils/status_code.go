package utils

import (
	"errors"
	"net/http"

	"passmanager/config"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetStatusCode(err error) int {
	statusCode := http.StatusInternalServerError
	if errors.Is(err, config.ErrRecordNotFound) {
		statusCode = http.StatusNotFound
	} else if errors.Is(err, config.ErrParameterMissing) {
		statusCode = http.StatusBadRequest
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		statusCode = http.StatusNotFound
	} else if errors.Is(err, config.ErrWrongPayload) {
		statusCode = http.StatusBadRequest
	} else if errors.Is(err, config.ErrDuplicateCustomer) {
		statusCode = http.StatusConflict
	}
	return statusCode
}
