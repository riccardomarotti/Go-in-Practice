package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

func main() {
	handler := NewHandler()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	manners.ListenAndServe(":8080", handler)

}

type handler struct{}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}

	fmt.Fprint(res, "Hello, my name is ", name)
}

func NewHandler() *handler {
	return &handler{}
}

func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	fmt.Println("Shutting down...")
	manners.Close()
	fmt.Println("Done")
}
