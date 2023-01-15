package main

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	ID          string `json:"id"`
	Name        string `json:"name"`
	IsCompleted bool   `json:"isCompleted"`
}

type CreateTodo struct {
	Name string `json:"name"`
	ID   string
}

type TodoList struct {
	DB    *gorm.DB
	todos []Todo
}

//func NewTodoList() TodoList {
//	return TodoList{
//		todos: make([]Todo, 0),
//	}
//}

func (tm *TodoList) GetAll() []Todo {
	tm.DB.Find(&tm.todos)
	return tm.todos
}

func (tm *TodoList) Create(createTodo CreateTodo) Todo {

	newTodo := Todo{
		ID:          createTodo.ID,
		Name:        createTodo.Name,
		IsCompleted: false,
	}

	//tm.todos = append(tm.todos, newTodo)
	tm.DB.Create(&newTodo)
	return newTodo
}
