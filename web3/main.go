package main

import (
	"fmt"
	"net/http"

	"example3.com/myapp"
)

func main() {
	fmt.Println("server start in port: 3000....")
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
