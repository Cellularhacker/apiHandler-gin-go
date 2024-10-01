package apiHandler

import (
	"fmt"
	"github.com/Cellularhacker/apiError-go"
	"github.com/Cellularhacker/logger-go"
	"github.com/Cellularhacker/util-go/pageInfo"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"strings"
)

type errorResp struct {
	Message string
	Success bool `json:"success"`
	Code    int  `json:"code"`
}

func JSONHandler(w http.ResponseWriter, r *http.Request, data interface{}, pi *pageInfo.Response, code int) {
	if err := SendDataResponse(w, data, pi, code); err != nil {
		ErrorHandler(w, r, err)
		return
	}
}

// ErrorHandler returns a httprouter handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, err *apiError.Error) {
	if err != nil {
		errMsg := fmt.Sprintf("%s : %s : %s", r.RequestURI, err.Error, err.Message)
		if err.Code >= 500 && err.Code < 600 {
			logger.L.Error(errMsg)
		} else if err.Code >= 400 && err.Code < 500 {
			logger.L.Warn(errMsg)
		}

		param := &errorResp{"", false, err.Code}
		if err.Code != apiError.UnknownServerError {
			param.Message = err.Message
		}

		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(strings.ToLower(err.Message), "invalid") {
			w.WriteHeader(http.StatusBadRequest)
		} else if err.Code >= 400 && err.Code < 600 {
			w.WriteHeader(err.Code)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		_ = json.NewEncoder(w).Encode(param)
	}
}

// APIHandler a string making it easy to handle errors
type APIHandler func(http.ResponseWriter, *http.Request, gin.Params) *apiError.Error

// Handle returns a httprouter handler
func (h APIHandler) Handle() gin.HandlerFunc {
	return func(context *gin.Context) {
		w := context.Writer
		r := context.Request
		p := context.Params
		if err := h(w, r, p); err != nil {
			errMsg := fmt.Sprintf("%s : %s : %s", r.RequestURI, err.Error, err.Message)
			if err.Code >= 500 && err.Code < 600 {
				logger.L.Error(errMsg)
			} else {
				logger.L.Warn(errMsg)
			}

			param := &errorResp{"", false, err.Code}
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

func API(handler APIHandler) gin.HandlerFunc {
	return handler.Handle()
}
