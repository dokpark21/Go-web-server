package app

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"time"
	"strconv"
)

// json을 반환할 것이기 때문에 render를 만들어준다.
var rd *render.Render

// DB를 따로 다루지 않고 web-server 안에서 in-memory로 데이터를 저장한다.
type Todo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Todo{}
	for _,v := range todoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}	

func addTestTodos() {
	todoMap[1] = &Todo{1, "Buy a milk1", false, time.Now()}
	todoMap[2] = &Todo{2, "Buy a milk2", true, time.Now()}
	todoMap[3] = &Todo{3, "Buy a milk3", false, time.Now()}
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	id := len(todoMap)+1
	todo := &Todo{id, name, false, time.Now()}
	todoMap[id] = todo
	rd.JSON(w, http.StatusOK, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok{
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})
	}else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	if todo, ok := todoMap[id]; ok{
		todo.Completed = complete
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func MakeHandler() http.Handler {
	todoMap = make(map[int]*Todo)
	rd = render.New()
	r := mux.NewRouter()

	addTestTodos()

	r.HandleFunc("/",indexHandler)
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")	
	r.HandleFunc("/todos", addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/todo-complete/{id:[0-9]+}", completeTodoHandler).Methods("GET")
	return r
}