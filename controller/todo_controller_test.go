package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-app/repository"
)

type mockTodoList struct {
	mock.Mock
}

func (m *mockTodoList) Create(createTodo repository.CreateTodo) repository.Todo {
	return repository.Todo{}
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

	//tc := TodoController{TodoRepo: repository.TodoList{}}
	//c := echo.New().NewContext(nil, nil)
	tc := TodoController{TodoRepo: mockTodoList}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := echo.New().NewContext(req, rec)
	err := tc.GetAllTodos(c)
	if err != nil {
		t.Error("Error occurred during test")
	}

	mockTodoList.AssertExpectations(t)

}
