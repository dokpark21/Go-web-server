package main

import (
	"os"
	// text, html 모두 template 패키지를 사용할 수 있다. text 패키지를 사용하게 되면 특수문자 탈락없이 사용할 수 있다.
	"html/template"
)

// template : html을 렌더링하는데 사용되는 템플릿 객체를 생성한다. 변경되지 않는 부분들을 템플릿화 하고 변경되는 부분들을 채워넣는다.

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "Park", Email: "psh03225@naver.com", Age: 25}
	user2 := User{Name: "Kim", Email: "kim1234@naver.com", Age: 40}
	users := []User{user, user2}
	// 빈공간에 채워넣을 수 있도록 템플릿을 생성한다.
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}
	// 템플릿에 내용을 채워넣어서 출력한다.
	// 템플릿안에 템플릿을 만들 수 있다. tmpl2.tmpl에 {{template "tmpl1.tmpl" .}}를 추가하면 tmpl1.tmpl의 내용이 출력된다.
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user2)
}
