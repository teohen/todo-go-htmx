package todo

import (
	"errors"

	"github.com/google/uuid"
	"github.com/teohen/todo-go-htmx/internal/domain"
)

type Repository interface {
	Save(todo domain.Todo) error
	GetAll() ([]domain.Todo, error)
	Get(id uuid.UUID) (domain.Todo, error)
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, todo domain.Todo) error
}

type TodoRepository struct {
	todos []domain.Todo
}

func NewTodoRepository() TodoRepository {
	return TodoRepository{}
}

func (TR *TodoRepository) Save(todo domain.Todo) error {
	TR.todos = append(TR.todos, todo)
	return nil
}

func (TR *TodoRepository) GetAll() ([]domain.Todo, error) {
	return TR.todos, nil
}

func (TR *TodoRepository) Get(id uuid.UUID) (domain.Todo, error) {
	for _, todo := range TR.todos {
		if todo.Id == id {
			return todo, nil
		}
	}

	return domain.Todo{}, nil
}

func (TR *TodoRepository) GetIndex(id uuid.UUID) int {
	for i, todo := range TR.todos {
		if todo.Id == id {
			return i
		}
	}

	return -1
}

func (TR *TodoRepository) Delete(id uuid.UUID) error {
	index := TR.GetIndex(id)
	if index < 0 {
		return errors.New("todo not found")
	}

	TR.todos = append(TR.todos[:index], TR.todos[index+1:]...)
	return nil
}

func (TR *TodoRepository) Update(id uuid.UUID, todo domain.Todo) error {
	index := TR.GetIndex(id)
	if index < 0 {
		return errors.New("todo not found")
	}

	TR.todos[index] = todo
	return nil
}
