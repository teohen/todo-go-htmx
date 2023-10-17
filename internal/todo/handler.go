package todo

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/teohen/todo-go-htmx/internal/domain"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
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
	err := r.ParseForm()

	if err != nil {
		fmt.Println("Error parsing the request form")
		w.WriteHeader(400)
		return
	}

	data := r.Form.Get("todo")

	if err != nil {
		fmt.Println("Error", err)
	}

	todo := domain.Todo{
		Title:   data,
		Checked: false,
	}

	_ = th.service.Save(todo)
	th.renderList(w, r)
}

func (th *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println("Error parsing form", err)
	}

	idTodoForm := chi.URLParam(r, "id")

	if err != nil {
		fmt.Println("Error", err)
	}

	idTodo, err := uuid.Parse(idTodoForm)

	if err != nil {
		fmt.Println("ERROR PARSING TODO ID FROM CLIENT", err)
	}

	err = th.service.Delete(idTodo)

	if err != nil {
		fmt.Println("Error deleting item", err)
	}
	th.renderList(w, r)
}

func (th *TodoHandler) Put(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println("Error parsing form", err)
	}

	idTodoForm := r.PostFormValue("id")

	if err != nil {
		fmt.Println("Error", err)
	}

	idTodo, err := uuid.Parse(idTodoForm)

	if err != nil {
		fmt.Println("ERROR PARSING TODO ID FROM CLIENT", err)
	}

	updatedTodo, err := th.service.Get(idTodo)

	if err != nil {
		fmt.Println("Error getting item", err)
	}

	updatedTodo.Checked = !updatedTodo.Checked

	err = th.service.Update(idTodo, updatedTodo)

	if err != nil {
		fmt.Println("Error updating todo", err)
	}
	th.renderList(w, r)
}

func (th *TodoHandler) renderList(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/list.html")

	if err != nil {
		fmt.Println("Error", err)
	}

	todoList := th.service.findAll()

	err = tmpl.ExecuteTemplate(w, "list.html", todoList)
	if err != nil {
		fmt.Println("error", err)
	}
}
