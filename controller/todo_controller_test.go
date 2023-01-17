package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
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

func TestGetAllTodos(t *testing.T) {
	expectedTodos := []repository.Todo{
		{ID: "1", Name: "Todo 1", IsCompleted: false},
		{ID: "2", Name: "Todo 2", IsCompleted: true},
	}

	mockTodoList := new(mockTodoList)
	mockTodoList.On("GetAll").Return(expectedTodos)

	tc := TodoController{TodoRepo: repository.TodoList{}}
	c := echo.New().NewContext(nil, nil)

	err := tc.GetAllTodos(c)
	if err != nil {
		t.Error("Error occurred during test")
	}

	mockTodoList.AssertExpectations(t)

}
