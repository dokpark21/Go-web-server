package app

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"strconv"
	"go-Todo-web.example/model"
)

// json을 반환할 것이기 때문에 render를 만들어준다.
var rd *render.Render

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	// list := []*model.Todo{}
	// for _,v := range todoMap {
	// 	list = append(list, v)
	// }
	list := model.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}	

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	// id := len(todoMap)+1
	// todo := &Todo{id, name, false, time.Now()}
	// todoMap[id] = todo
	todo := model.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := model.RemoveTodo(id)

	if ok {
			rd.JSON(w, http.StatusOK, Success{true})	
	}else {
			rd.JSON(w, http.StatusOK, Success{false})
	}

	// if _, ok := todoMap[id]; ok{
	// 	delete(todoMap, id)
	// 	rd.JSON(w, http.StatusOK, Success{true})
	// }else {
	// 	rd.JSON(w, http.StatusOK, Success{false})
	// }
}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"

	ok := model.CompleteTodo(id, complete)

	if ok {
			rd.JSON(w, http.StatusOK, Success{true})	
	}else {
			rd.JSON(w, http.StatusOK, Success{false})
	}
	// if todo, ok := todoMap[id]; ok{
	// 	todo.Completed = complete
	// 	rd.JSON(w, http.StatusOK, Success{true})
	// } else {
	// 	rd.JSON(w, http.StatusOK, Success{false})
	// }
}

func MakeHandler() http.Handler {
	rd = render.New()
	r := mux.NewRouter()

	r.HandleFunc("/",indexHandler)
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")	
	r.HandleFunc("/todos", addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/todo-complete/{id:[0-9]+}", completeTodoHandler).Methods("GET")
	return r
}