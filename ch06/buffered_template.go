package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseFiles("simple.html"))

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin' da castle.",
	}

	var b bytes.Buffer
	err := t.Execute(w, p)
	if err != nil {
		fmt.Fprint(w, "An error occurred...")
		return
	}

	b.WriteTo(w)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
