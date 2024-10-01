package apiHandler

import (
	"github.com/Cellularhacker/apiError-go"
	"net/http"
)

func SendRedirectResponse(w http.ResponseWriter, location string, statusCode int) *apiError.Error {
	w.WriteHeader(statusCode)
	w.Header().Set("Location", location)

	return nil
}
