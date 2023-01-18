package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"todo-app/controller"
	"todo-app/repository"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&repository.Todo{})

	//todoController := &controller.TodoController{TodoRepo: repository.TodoList{DB: db}}
	todoController := &controller.TodoController{TodoRepo: &repository.TodoList{DB: db}}

	e := echo.New()

	e.GET("/todos", todoController.GetAllTodos)
	e.POST("/todos/create", todoController.CreateTodo)

	e.Start(":8888")
}
