package model

import (
	"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // 패키지를 암시적으로 사용(명시적X)
)

type sqliteHandler struct {
	
}

// memoryHandler의 메서드 정의
func (s *sqliteHandler)GetTodos() []*Todo{
	return nil
}

func (s *sqliteHandler)AddTodo(name string) *Todo{
	return nil
}

func (s *sqliteHandler)RemoveTodo(id int) bool{
	return false
}

func (s *sqliteHandler)CompleteTodo(id int, complete bool) bool{
	return false
}

func newSqliteHandler() dbHandler{
	database, err := sql.Open("sqlite3", "./test.db")
	return &sqliteHandler{}
}