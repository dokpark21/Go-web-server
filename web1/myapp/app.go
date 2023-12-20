package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	// 결과를 name이란 변수에 집어 넣는다
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	// Body에서 JSON 형태에 데이터를 읽어서 user 변수에 Decode한다.
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}

	user.CreatedAt = time.Now()
	// go structure 형태의 데이터를 다시 Marshal해줘서 다시 JSON형태의 데이터로 바꾼다
	data, _ := json.Marshal(user)
	// header에 content type을 application/json이라고 명시해준다
	w.Header().Add("content-type", "application/json")
	// header에 올바른 상태라고 명시
	w.WriteHeader(http.StatusOK)
	// 결과값을 다시 client에 전송
	fmt.Fprint(w, string(data))
}

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	return mux
}
