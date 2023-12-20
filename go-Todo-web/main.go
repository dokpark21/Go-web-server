package main

import (
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/choiseungyoun/go-Todo-web/app"
)

func main() {
	m := app.MakeHandler()
	n := negroni.Classic()
	n.useHandler(m)

	http.ListenAndServe(":3000", n)
}