package apiHandler

import (
	"fmt"
	"github.com/Cellularhacker/apiError-go"
	"github.com/Cellularhacker/logger-go"
	"github.com/goccy/go-json"
	"net/http"
	"strings"
)

// APIHandler a string making it easy to handle errors
type APIHandler func(http.ResponseWriter, *http.Request) *apiError.Error

// Handle returns a httprouter handler
func (h APIHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			errMsg := fmt.Sprintf("%s : %s : %s", r.RequestURI, err.Error, err.Message)
			if err.Code >= 500 && err.Code < 600 {
				logger.L.Error(errMsg)
			} else {
				logger.L.Warn(errMsg)
			}

			param := &ErrorResp{"", false, err.Code}
			if err.Code != apiError.UnknownServerError {
				param.Message = err.Message
			}

			w.Header().Set("Content-Type", "application/json")
			if strings.Contains(strings.ToLower(err.Message), "invalid") {
				w.WriteHeader(http.StatusBadRequest)
			} else if err.Code >= 300 && err.Code < 600 {
				w.WriteHeader(err.Code)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			_ = json.NewEncoder(w).Encode(param)
		}
	}
}

func Handle(handler APIHandler) http.HandlerFunc {
	return handler.Handle()
}
