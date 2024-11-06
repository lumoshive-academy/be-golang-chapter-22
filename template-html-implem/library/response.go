package library

import (
	"be-golang-chapter-22/template-html-implem/model"
	"encoding/json"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, message string, data any) {
	badResponse := model.Response{
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
	}
	json.NewEncoder(w).Encode(badResponse)

}

func BadResponse(w http.ResponseWriter, message string) {
	badResponse := model.Response{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Data:       nil,
	}
	json.NewEncoder(w).Encode(badResponse)
}
