package app

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"time"
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

func MakeHandler() http.Handler {
	todoMap = make(map[int]*Todo)
	rd = render.New()
	r := mux.NewRouter()

	addTestTodos()

	r.HandleFunc("/",indexHandler)
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")	
	return r
}