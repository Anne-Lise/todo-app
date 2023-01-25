package repository

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Name        string `json:"name"`
	IsCompleted bool   `json:"isCompleted"`
}

type CreateTodo struct {
	gorm.Model
	Name string `json:"name"`
	//ID   string

}

type TodoList struct {
	DB    *gorm.DB
	todos []Todo
}
type TodoRepository interface {
	Create(createTodo CreateTodo) Todo
	GetAll() []Todo
}

func (tm *TodoList) GetAll() []Todo {
	tm.DB.Find(&tm.todos)
	return tm.todos
}

func (tm *TodoList) Create(createTodo CreateTodo) Todo {

	newTodo := Todo{
		Name: createTodo.Name,
	}

	tm.DB.Create(&newTodo)
	return newTodo
}
