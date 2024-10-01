package handler

import (
	"fmt"
	"github.com/Cellularhacker/apiError-go"
	"github.com/Cellularhacker/apiHandler-go"
	"github.com/Cellularhacker/logger-go"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"strings"
)

// Gin a string making it easy to handle errors
type Gin func(http.ResponseWriter, *http.Request, gin.Params) *apiError.Error

// Handle returns a httprouter handler
func (gh Gin) Handle() gin.HandlerFunc {
	return func(context *gin.Context) {
		w := context.Writer
		r := context.Request
		p := context.Params
		if err := gh(w, r, p); err != nil {
			errMsg := fmt.Sprintf("%s : %s : %s", r.RequestURI, err.Error, err.Message)
			if err.Code >= 500 && err.Code < 600 {
				logger.L.Error(errMsg)
			} else {
				logger.L.Warn(errMsg)
			}

			param := &apiHandler.ErrorResp{Code: err.Code}
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

func Handle(handler Gin) gin.HandlerFunc {
	return handler.Handle()
}

func ToGin(h apiHandler.APIHandler) Gin {
	return func(w http.ResponseWriter, r *http.Request, p gin.Params) *apiError.Error {
		return h(w, r)
	}
}
