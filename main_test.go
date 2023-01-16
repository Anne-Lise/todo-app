package main

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
	"todo-app/repository"
)

type mockTodoList struct {
	mock.Mock
}

func (m *mockTodoList) GetAll() []repository.Todo {
	args := m.Called()
	return args.Get(0).([]repository.Todo)
}

func (m *mockTodoList) Create(createTodo repository.CreateTodo) repository.Todo {
	args := m.Called(createTodo)
	return args.Get(0).(repository.Todo)
}

func TestGetAll(t *testing.T) {
	testTodos := []repository.Todo{
		{ID: "1", Name: "Todo 1", IsCompleted: false},
		{ID: "2", Name: "Todo 2", IsCompleted: true},
	}

	mockTodoList := new(mockTodoList)
	mockTodoList.On("GetAll").Return(testTodos)

	e := echo.New()

	e.GET("/todos", func(c echo.Context) error {
		todos := mockTodoList.GetAll()
		return c.JSON(http.StatusOK, todos)
	})
}

func TestCreate(t *testing.T) {
	testTodo := repository.Todo{ID: "1", Name: "Todo 1", IsCompleted: false}

	mockTodoList := new(mockTodoList)
	mockTodoList.On("Create", repository.CreateTodo{Name: "Todo 1"}).Return(testTodo)
	e := echo.New()

	e.POST("/todos/create", func(c echo.Context) error {
		requestBody := repository.CreateTodo{}
		err := c.Bind(&requestBody)
		if err != nil {
			return err
		}

		todo := mockTodoList.Create(requestBody)
		return c.JSON(http.StatusCreated, todo)
	})
}
