package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"go-Todo-web.example/app"
)

func main() {
	m := app.MakeHandler()
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
