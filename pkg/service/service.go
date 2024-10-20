package service

import (
	"github.com/Timur000111/todo-app"
	"github.com/Timur000111/todo-app/pkg/repository"
)

type Autthorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
}

type TodoItem interface {
}

type Service struct {
	Autthorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autthorization: NewAuthService(repos.Autthorization),
		TodoList:       NewTodoListService(repos.TodoList),
	}
}
