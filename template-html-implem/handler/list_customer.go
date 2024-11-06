package handler

import (
	"be-golang-chapter-22/template-html-implem/database"
	"be-golang-chapter-22/template-html-implem/model"
	"be-golang-chapter-22/template-html-implem/repository"
	"be-golang-chapter-22/template-html-implem/service"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

var templates = template.Must(template.ParseGlob("view/*.html"))

func ListCustomer(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB()

	if err != nil {
		fmt.Println("err ", err)
	}

	repo := repository.NewCustomerRepository(db)
	serviceCustomer := service.NewCustomerService(repo)

	customers, err := serviceCustomer.AllCustomer()
	if err != nil {
		fmt.Println("err ", err)
	}

	// fmt.Println("data :", *customers)
	templates.ExecuteTemplate(w, "list-user-view", *customers)
}

func FormRegist(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "register-view", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var name, email, username, password string
	if r.Method == http.MethodPost {
		name = r.FormValue("name")
		email = r.FormValue("email")
		username = r.FormValue("username")
		password = r.FormValue("password")
	}

	fmt.Println("data :", name, email, username, password)

	db, err := database.InitDB()
	if err != nil {
		fmt.Println("err ", err)
	}

	customerRegister := model.Customer{
		Name:     name,
		Username: username,
		Password: password,
		Email:    email,
	}

	token := uuid.New()
	customerRegister.Status = "active"
	customerRegister.Token = token.String()

	repo := repository.NewCustomerRepository(db)
	serviceCustomer := service.NewCustomerService(repo)

	err = serviceCustomer.Register(&customerRegister)
	if err != nil {
		fmt.Println("Error :", err)
	}
	http.Redirect(w, r, "/all-customer", http.StatusSeeOther)
}

func CustomerDetail(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	id_int, _ := strconv.Atoi(id)

	db, err := database.InitDB()

	if err != nil {
		fmt.Println("err ", err)
	}

	repo := repository.NewCustomerRepository(db)
	serviceCustomer := service.NewCustomerService(repo)

	customers, err := serviceCustomer.CustomerByID(id_int)
	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Println(customers)
	templates.ExecuteTemplate(w, "user-detail-view", *customers)
}
