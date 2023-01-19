package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestTodoList_GetAll(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open("postgres", db)
	defer gormDB.Close()
	tm := TodoList{DB: gormDB}

	todos := []Todo{
		{ID: "1", Name: "Todo 1", IsCompleted: false},
		{ID: "2", Name: "Todo 2", IsCompleted: true},
	}
	mock.ExpectQuery("SELECT * FROM \"todos\"").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "is_completed"}).AddRow(1, "Todo 1", false).AddRow(2, "Todo 2", true))
	result := tm.GetAll()
	assert.Equal(t, todos, result)
}
