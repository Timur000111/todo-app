package repository

import (
	"github.com/Timur000111/todo-app"
	"github.com/jmoiron/sqlx"
)

type Autthorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)

}

type TodoItem interface {
}

type Repository struct {
	Autthorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autthorization: NewAuthPostgres(db),
		TodoList:       NewTodoListPostgres(db),
	}
}
