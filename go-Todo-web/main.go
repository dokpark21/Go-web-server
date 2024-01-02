package main

import (
	"net/http"

	"go-Todo-web.example/app"
)

func main() {

	m := app.MakeHandler("./test.db")
	defer m.Close()

	err := http.ListenAndServe(":3000", m)
	if err != nil {
		panic(err)
	}
}
