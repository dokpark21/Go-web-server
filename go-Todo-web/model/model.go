package model

import (
	"time"
)

// DB를 따로 다루지 않고 web-server 안에서 in-memory로 데이터를 저장한다.
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// db handler 정의
type DBHandler interface {
	GetTodos() []*Todo
	AddTodo(name string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	// db 사용 생명주기 관리: db 연결 종료를 사용하는 쪽에 위임(sqliteHandler의 Close 함수 사용)
	Close()
}

// handler는 dbHandler의 instance를 들고 있다.
var handler DBHandler

func NewDBHandler() DBHandler {
	return newSqliteHandler()

}
