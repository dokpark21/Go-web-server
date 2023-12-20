package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:  "John",
		Email: "John@naver.com",
	}

	// render 패키지를 이용한 JSON 응답(아래 4줄을 한줄로 줄일 수 있음)
	rd.JSON(w, http.StatusOK, user)

	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}

	user.CreatedAt = time.Now()

	rd.JSON(w, http.StatusOK, user)
	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// render 패키지를 사용해서 템플릿을 불러오려면 반드시 templates 폴더 안에 확장자명 .tmpl을 사용해야 한다.
	user := User{
		Name:  "John",
		Email: "John@naver.com",
	}
	rd.HTML(w, http.StatusOK, "body", user)
}

func main() {
	// option을 통해 탬플릿 폴더명이나 확장자명을 지정할 수 있다.
	rd = render.New(
		render.Options{
			Extensions: []string{".html", ".tmpl"},
			Directory:  "templates", // 다른 폴더명으로 설정 가능
			Layout:     "hello",
		},
	)
	// gorilla/pat 패키지를 이용한 라우터 생성
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	// 자동으로 public 폴더를 등록해주는 미들웨어
	// logger, recovery 미들웨어를 사용하기 위해 negroni 패키지를 사용
	n := negroni.Classic()
	// n에 mux를 등록(wrapping)
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
