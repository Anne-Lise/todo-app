package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockTodoList struct {
	mock.Mock
}

func (m *mockTodoList) GetAll() []Todo {
	args := m.Called()
	return args.Get(0).([]Todo)
}

func (m *mockTodoList) Create(createTodo CreateTodo) Todo {
	args := m.Called(createTodo)
	return args.Get(0).(Todo)
}

func TestGetAll(t *testing.T) {
	testTodos := []Todo{
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
	testTodo := Todo{ID: "1", Name: "Todo 1", IsCompleted: false}

	mockTodoList := new(mockTodoList)
	mockTodoList.On("Create", CreateTodo{Name: "Todo 1"}).Return(testTodo)
	e := echo.New()

	e.POST("/todos/create", func(c echo.Context) error {
		requestBody := CreateTodo{}
		err := c.Bind(&requestBody)
		if err != nil {
			return err
		}

		todo := mockTodoList.Create(requestBody)
		return c.JSON(http.StatusCreated, todo)
	})
}
