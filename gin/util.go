package handler

import (
	"github.com/Cellularhacker/apiHandler-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CoffeeGET(w http.ResponseWriter, _ *http.Request, _ gin.Params) Gin {
	return ToGin(apiHandler.CoffeeGET)
}

func CreateTestPOST(w http.ResponseWriter, _ *http.Request, _ gin.Params) Gin {
	return ToGin(apiHandler.CreateTestPOST)
}

func OKTestPOST(w http.ResponseWriter, _ *http.Request, _ gin.Params) Gin {
	return ToGin(apiHandler.OKTestPOST)
}

func Ping(_ http.ResponseWriter, _ *http.Request, _ gin.Params) Gin {
	return ToGin(apiHandler.Ping)
}

func Info() Gin {
	return ToGin(apiHandler.Info)
}
