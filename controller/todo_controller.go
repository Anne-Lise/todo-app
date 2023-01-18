package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"todo-app/repository"
)

type TodoController struct {
	//TodoRepo repository.TodoList
	TodoRepo repository.TodoRepository
}

func (tc *TodoController) CreateTodo(c echo.Context) error {
	requestBody := repository.CreateTodo{}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	todo := tc.TodoRepo.Create(requestBody)
	return c.JSON(http.StatusCreated, todo)
}

func (tc *TodoController) GetAllTodos(c echo.Context) error {
	todos := tc.TodoRepo.GetAll()

	fmt.Print("todos:", todos)

	return c.JSON(http.StatusOK, todos)
}
