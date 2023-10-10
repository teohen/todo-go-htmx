package todo

import (
	"fmt"
	"net/http"
	"text/template"
)

type todoHandler struct {
	service Service
}

func NewTodoHandler(service Service) todoHandler {
	return todoHandler{
		service: service,
	}
}

func (th *todoHandler) HandleGet(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("ola")
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		fmt.Println("Error", err)
	}

	todoList := th.service.Get()

	data := todoList[0]

	fmt.Println(data)

	err = tmpl.ExecuteTemplate(w, "index.html", data)

	if err != nil {
		fmt.Println("error", err)
	}
}
