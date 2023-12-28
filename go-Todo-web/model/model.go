package model

import (
	"time"
)

// DB를 따로 다루지 않고 web-server 안에서 in-memory로 데이터를 저장한다.
type Todo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

//db handler 정의 
type dbHandler interface {
	GetTodos() []*Todo
	AddTodo(name string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
}

// handler는 dbHandler의 instance를 들고 있다.
var handler dbHandler

func init() {
	// handler = newMemoryHandler()
	handler = newSqliteHandler()
}

func GetTodos() []*Todo {
	return handler.GetTodos()
}

func AddTodo(name string) *Todo {
	return handler.AddTodo(name)
}

func RemoveTodo(id int) bool {
	return handler.RemoveTodo(id)
}

func CompleteTodo(id int, complete bool) bool {
	return handler.CompleteTodo(id, complete)
}