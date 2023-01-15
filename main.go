package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Todo{})
	//tm := NewTodoList()
	tm := &TodoList{DB: db}
	//tm := NewTodoList()
	e := echo.New()

	e.GET("/todos", func(c echo.Context) error {
		todos := tm.GetAll()

		return c.JSON(http.StatusOK, todos)
	})

	e.POST("todos/create", func(c echo.Context) error {
		requestBody := CreateTodo{}

		err := c.Bind(&requestBody)
		if err != nil {
			return err
		}

		todo := tm.Create(requestBody)
		return c.JSON(http.StatusCreated, todo)
	})

	e.Start(":8888")
}
