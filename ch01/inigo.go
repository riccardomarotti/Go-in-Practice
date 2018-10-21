package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is Inigo Montoya")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Listening on port localhost:4000...")
	http.ListenAndServe("localhost:4000", nil)
}
