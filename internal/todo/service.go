package todo

import (
	"fmt"

	"github.com/teohen/todo-go-htmx/internal/domain"
)

type todoService struct {
	db string
}

type Service interface {
	Get() []domain.Todo
}

func NewTodoService() Service {
	return &todoService{
		db: "banco",
	}
}

func (db *todoService) Get() []domain.Todo {
	var todos []domain.Todo
	fmt.Println("get todos")
	todos = append(todos, domain.Todo{
		Title:   "TITULO",
		Checked: false,
	})

	return todos
}
