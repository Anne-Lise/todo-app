package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"regexp"
	"testing"
)

func TestTodoList_GetAll(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open("postgres", db)
	defer gormDB.Close()
	tm := TodoList{DB: gormDB}

	todos := []Todo{
		{Model: gorm.Model{ID: 1}, Name: "Todo 1", IsCompleted: false},
		{Model: gorm.Model{ID: 2}, Name: "Todo 2", IsCompleted: true},
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"todos\"")).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "is_completed"}).AddRow(1, "Todo 1", false).AddRow(2, "Todo 2", true))
	result := tm.GetAll()
	assert.Nil(t, mock.ExpectationsWereMet())
	assert.Equal(t, todos, result)
}

func TestTodoList_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open("postgres", db)
	defer gormDB.Close()
	tm := TodoList{DB: gormDB}

	createTodo := CreateTodo{Name: "Todo 1"}
	//newTodo := Todo{Model: gorm.Model{}, Name: "Todo 1", IsCompleted: false}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `todos` (`created_at`,`updated_at`,`deleted_at`,`name`,`is_completed`) VALUES (?,?,?,?,?)")).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Todo 1", false).
		WillReturnResult(sqlmock.NewResult(1, 1))
	result := tm.Create(createTodo)
	fmt.Printf("Result: created_at: %s,updated_at: %s,deleted_at: %s,ID: %d, Name: %s, IsCompleted: %t\n", result.CreatedAt, result.UpdatedAt, result.DeletedAt, result.ID, result.Name, result.IsCompleted)
	//fmt.Printf("Expected: created_at: %s,updated_at: %s,deleted_at: %s,ID: %d, Name: %s, IsCompleted: %t\n", newTodo.CreatedAt, newTodo.UpdatedAt, newTodo.DeletedAt, newTodo.ID, newTodo.Name, newTodo.IsCompleted)
	assert.Nil(t, mock.ExpectationsWereMet())
	//assert.Equal(t, newTodo, result)
	//if err := mock.ExpectationsWereMet(); err != nil {
	//	t.Errorf("there were unfulfilled expectations: %s", err)
	//}
}
