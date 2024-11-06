package middleware

import (
	"be-golang-chapter-22/template-html-implem/database"
	"be-golang-chapter-22/template-html-implem/model"
	"be-golang-chapter-22/template-html-implem/repository"
	"be-golang-chapter-22/template-html-implem/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("token")

		db, err := database.InitDB()
		if err != nil {
			fmt.Println("err ", err)
		}
		repo := repository.NewCustomerRepository(db)
		serviceCustomer := service.NewCustomerService(repo)

		token := serviceCustomer.CheckToken(authHeader)
		if token == "" {
			badResponse := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       nil,
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(badResponse)
			return
		}

		// Melanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
