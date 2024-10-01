package handler

import (
	"github.com/Cellularhacker/apiHandler-go"
)

func CoffeeGET() Gin {
	return ToGin(apiHandler.CoffeeGET)
}

func CreateTestPOST() Gin {
	return ToGin(apiHandler.CreateTestPOST)
}

func OKTestPOST() Gin {
	return ToGin(apiHandler.OKTestPOST)
}

func Ping() Gin {
	return ToGin(apiHandler.Ping)
}

func Info() Gin {
	return ToGin(apiHandler.Info)
}
