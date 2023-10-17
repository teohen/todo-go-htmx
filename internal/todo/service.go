package todo

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/teohen/todo-go-htmx/internal/domain"
)

type TodoService struct {
	repository Repository
}

type Service interface {
	Save(todo domain.Todo) domain.Todo
	Get(id uuid.UUID) (domain.Todo, error)
	findAll() []domain.Todo
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, todo domain.Todo) error
}

func NewTodoService(db Repository) Service {
	return &TodoService{
		repository: db,
	}
}

func (TS *TodoService) findAll() []domain.Todo {
	var todos []domain.Todo
	todos, err := TS.repository.GetAll()

	if err != nil {
		fmt.Println("ERRO FINDING ALL TODOS", err)
		return todos
	}

	return todos
}

func (TS *TodoService) Get(id uuid.UUID) (domain.Todo, error) {
	var todo domain.Todo
	todo, err := TS.repository.Get(id)

	if err != nil {
		fmt.Println("ERROR GETTING ONE TODO: ", err)
		return todo, err
	}

	return todo, nil
}

func (TS *TodoService) Save(todo domain.Todo) domain.Todo {
	todo.Id = uuid.New()

	err := TS.repository.Save(todo)

	if err != nil {
		fmt.Println("ERROR SAVING ONE TODO: ", err)
		return domain.Todo{}
	}

	return todo
}

func (TS *TodoService) Delete(id uuid.UUID) error {
	err := TS.repository.Delete(id)

	if err != nil {
		fmt.Println("ERROR DELETING ONE TODO: ", err)
		return err
	}

	return nil
}

func (TS *TodoService) Update(id uuid.UUID, todo domain.Todo) error {
	err := TS.repository.Update(id, todo)

	if err != nil {
		fmt.Println("ERROR UPDATING ONE TODO:", err)
		return err
	}

	return nil
}
