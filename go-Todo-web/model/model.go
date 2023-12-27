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

var todoMap map[int]*Todo

func GetTodos() []*Todo {
	return nil
}

func AddTodo(name string) *Todo {
	return nil
}