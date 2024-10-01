package handler

import (
	"github.com/Cellularhacker/apiHandler-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CoffeeGET(w http.ResponseWriter, _ *http.Request, _ gin.Params) gin.HandlerFunc {
	return ToGin(apiHandler.CoffeeGET).Handle()
}

func CreateTestPOST(w http.ResponseWriter, _ *http.Request, _ gin.Params) gin.HandlerFunc {
	return ToGin(apiHandler.CreateTestPOST).Handle()
}

func OKTestPOST(w http.ResponseWriter, _ *http.Request, _ gin.Params) gin.HandlerFunc {
	return ToGin(apiHandler.OKTestPOST).Handle()
}

func Ping(_ http.ResponseWriter, _ *http.Request, _ gin.Params) gin.HandlerFunc {
	return ToGin(apiHandler.Ping)
}

func Info() gin.HandlerFunc {
	return ToGin(apiHandler.Info).Handle()
}
