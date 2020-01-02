package main

import (
	"github.com/FelipeGalante/go-basic-projects/hello-world/hello"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", servePage)
	http.ListenAndServe(":8080", nil)
}

func servePage(writer http.ResponseWriter, reqest *http.Request) {
	io.WriteString(writer, hello.Hello())
}
