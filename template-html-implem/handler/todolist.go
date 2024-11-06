package handler

import (
	"be-golang-chapter-22/template-html-implem/library"
	"net/http"
)

func AllTodoList(w http.ResponseWriter, r *http.Request) {

	// get data todo list dari database
	library.SuccessResponse(w, "data todolist", nil)
}
