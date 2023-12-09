package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/tomsolem/devcontainer/hello"
)

type Hello struct{}

func handle(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, hello.Hello())
}

func main() {
	portNumber := "9000"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
