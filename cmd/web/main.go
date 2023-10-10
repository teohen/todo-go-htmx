package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/teohen/todo-go-htmx/internal/todo"
)

func main() {
	fmt.Println("server running")

	router := chi.NewRouter()

	service := todo.NewTodoService()

	handler := todo.NewTodoHandler(service)

	router.Get("/", handler.HandleGet)

	http.ListenAndServe(":3000", router)

}
