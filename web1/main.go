package main

import (
	"log"
	"net/http"

	"webserver-practice.com/myapp"
)

func main() {

	log.Println("server is start port: 3000")
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
