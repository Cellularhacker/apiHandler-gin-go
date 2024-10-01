package apiHandler

import (
	"github.com/Cellularhacker/apiError-go"
	"github.com/Cellularhacker/util-go"
	"net/http"
	"strings"
	"time"
)

type InfoResp struct {
	ServerTimeUnix int64  `json:"server_time"`
	TimeZone       string `json:"time_zone"`
}

const (
	queryAllOK = "query_all_ok"
)

func CoffeeGET(w http.ResponseWriter, _ *http.Request, _ interface{}) *apiError.Error {
	return SendMsgResponse(w, "I reject to boil your teapot.", http.StatusTeapot)
}

func CreateTestPOST(w http.ResponseWriter, _ *http.Request, _ interface{}) *apiError.Error {
	return SendSuccessResponse(w, true)
}

func OKTestPOST(w http.ResponseWriter, _ *http.Request, _ interface{}) *apiError.Error {
	return SendSuccessResponse(w, false)
}

func Ping(w http.ResponseWriter, _ *http.Request, _ interface{}) *apiError.Error {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("pong"))
	if err != nil {
		return apiError.InternalServerError(err)
	}

	return nil
}

func Info(w http.ResponseWriter, _ *http.Request, _ interface{}) *apiError.Error {
	return SendDataResponse(w, getInfo(), nil, http.StatusOK)
}

func getContentType(r *http.Request) string {
	return strings.ToLower(r.Header.Get("Content-Type"))
}

func getInfo() *InfoResp {
	return &InfoResp{ServerTimeUnix: time.Now().Unix(), TimeZone: util.Loc.String()}
}
