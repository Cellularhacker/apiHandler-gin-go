package apiHandler

import (
	"bytes"
	"fmt"
	"github.com/Cellularhacker/apiError-go"
	"github.com/Cellularhacker/util-go/uTime"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func ExtractCSVHeaders(v interface{}) string {
	headers := make([]string, 0)

	val := reflect.ValueOf(v)
	for i := 0; i < val.Type().NumField(); i++ {
		h := val.Type().Field(i).Tag.Get("csv")
		if h == "" {
			h = val.Type().Field(i).Tag.Get("json")
		}

		headers = append(headers, h)
	}

	return fmt.Sprintf("%s\n", strings.Join(headers, ","))
}

func ExtractCSVData(v interface{}) string {
	data := make([]string, 0)

	val := reflect.ValueOf(v)
	for i := 0; i < val.NumField(); i++ {
		data = append(data, fmt.Sprintf("%v", val.Field(i).Interface()))
	}

	return fmt.Sprintf("%s\n", strings.Join(data, ","))
}

func SendCSVResponseFromString(w http.ResponseWriter, fileNamePrefix string, csvStr *string) *apiError.Error {
	br := bytes.NewBufferString(*csvStr)
	fileName := fmt.Sprintf("%s_%s.csv", fileNamePrefix, uTime.GetKSTDateStr(nil))

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Length", strconv.Itoa(len(*csvStr)))

	_, err := io.Copy(w, br)
	if err != nil {
		return apiError.InternalServerError(err)
	}

	return nil
}

func SendCSVResponse(w http.ResponseWriter, fileNamePrefix string, csv *bytes.Buffer) *apiError.Error {
	fileName := fmt.Sprintf("%s_%s.csv", fileNamePrefix, uTime.GetKSTDateStr(nil))

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Length", strconv.Itoa(csv.Len()))

	_, err := io.Copy(w, csv)
	if err != nil {
		return apiError.InternalServerError(err)
	}

	return nil
}

func SendCSVResponseWithFileName(w http.ResponseWriter, fileName string, csvStr *string) *apiError.Error {
	br := bytes.NewBufferString(*csvStr)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.csv", fileName))
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Length", strconv.Itoa(len(*csvStr)))

	_, err := io.Copy(w, br)
	if err != nil {
		return apiError.InternalServerError(err)
	}

	return nil
}
