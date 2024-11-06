package main

import (
	"be-golang-chapter-22/template-html-implem/handler"
	"be-golang-chapter-22/template-html-implem/middleware"
	"fmt"
	"net/http"
)

func main() {
	mainRoute := http.NewServeMux()

	r := http.NewServeMux()

	// route view
	r.HandleFunc("/all-customer", handler.ListCustomer)
	r.HandleFunc("/form-register", handler.FormRegist)

	// route action
	r.HandleFunc("/register", handler.Register)
	r.HandleFunc("/customre-detail/{id}", handler.CustomerDetail)

	// route api
	todoRoute := http.NewServeMux()
	todoRoute.HandleFunc("GET /all-todolist", handler.AllTodoList)
	middRoute := middleware.Middleware(todoRoute)

	mainRoute.Handle("/", r)
	mainRoute.Handle("/todolist/", http.StripPrefix("/todolist", middRoute))

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mainRoute)
}
