package handler

import (
	"be-golang-chapter-22/template-html-implem/library"
	"be-golang-chapter-22/template-html-implem/model"
	"be-golang-chapter-22/template-html-implem/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type CustomerHandler struct {
	CustomerService service.CustomerService
}

func NewCustomerHandler(cs service.CustomerService) CustomerHandler {
	return CustomerHandler{CustomerService: cs}
}

func (ch *CustomerHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	customer := model.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		library.BadResponse(w, err.Error())
		return
	}

	err = ch.CustomerService.LoginService(customer)
	if err != nil {
		library.BadResponse(w, "Account not found")
		return
	}

	library.SuccessResponse(w, "Login Success", customer)
}

func (ch *CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")

	id_int, _ := strconv.Atoi(id)

	customer, err := ch.CustomerService.CustomerByID(id_int)
	if err != nil {
		library.BadResponse(w, "Account not found")
		return
	}

	library.SuccessResponse(w, "Login Success", customer)
}
