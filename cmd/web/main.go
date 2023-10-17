package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/teohen/todo-go-htmx/internal/todo"
)

func main() {
	fmt.Println("server running")

	router := chi.NewRouter()

	repository := todo.NewTodoRepository()

	service := todo.NewTodoService(&repository)

	handler := todo.NewTodoHandler(service)

	router.Get("/", handler.GetAll)
	router.Post("/todo", handler.Create)
	router.Delete("/todo/{id}", handler.Delete)
	router.Put("/todo", handler.Put)

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		log.Fatal("Error listening to port", err)
	}
}
