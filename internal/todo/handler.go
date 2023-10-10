package todo

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/uuid"
	"github.com/teohen/todo-go-htmx/internal/domain"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type TodoHandler struct {
	service Service
}

func NewTodoHandler(service Service) Handler {
	return &TodoHandler{
		service: service,
	}
}

func (th *TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")

	if err != nil {
		fmt.Println("Error", err)
	}

	todoList := th.service.findAll()

	err = tmpl.ExecuteTemplate(w, "index.html", todoList)
	if err != nil {
		fmt.Println("error", err)
	}
}

func (th *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	data := r.Form.Get("todo")

	tmpl, err := template.ParseFiles("./templates/list.html")

	if err != nil {
		fmt.Println("Error", err)
	}

	todo := domain.Todo{
		Title:   data,
		Checked: false,
	}

	_ = th.service.Save(todo)

	todoList := th.service.findAll()

	err = tmpl.ExecuteTemplate(w, "list.html", todoList)

	if err != nil {
		fmt.Println("error", err)
	}
}

func (th *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	data := r.PostFormValue("id")

	tmpl, err := template.ParseFiles("./templates/list-item.html")

	if err != nil {
		fmt.Println("Error", err)
	}

	idTodo, err := uuid.Parse(data)

	if err != nil {
		fmt.Println("ERROR PARSING TODO ID FROM CLIENT", err)
	}

	err = th.service.Delete(idTodo)

	err = tmpl.ExecuteTemplate(w, "./templates/list-item.html", nil)

	if err != nil {
		fmt.Println("error", err)
	}
}
