package model

import (
	"time"
)

// memoryHandler 정의
type memoryHandler struct {
	todoMap map[int]*Todo
}

// memoryHandler의 메서드 정의
func (m *memoryHandler)GetTodos() []*Todo{
	list := []*Todo{}
	for _,v := range m.todoMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler)AddTodo(name string) *Todo{
	id := len(m.todoMap)+1
	todo := &Todo{id, name, false, time.Now()}
	m.todoMap[id] = todo
	return todo
}

func (m *memoryHandler)RemoveTodo(id int) bool{
	if _, ok := m.todoMap[id]; ok{
		delete(m.todoMap, id)
		return true
	}
	return false
}

func (m *memoryHandler)CompleteTodo(id int, complete bool) bool{
	if todo, ok := m.todoMap[id]; ok{
		todo.Completed = complete
		return true
	}
	return false
}

func newMemoryHandler() dbHandler{
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	return m
}