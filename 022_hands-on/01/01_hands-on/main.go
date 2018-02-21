package main

import (
	"fmt"
	"io"
	"net/http"
)

func logRequest(req *http.Request) {
	r := *req
	fmt.Println(r.Method, r.URL)
}

func index(res http.ResponseWriter, req *http.Request) {
	logRequest(req)
	io.WriteString(res, "index\n")
}

func dog(res http.ResponseWriter, req *http.Request) {
	logRequest(req)
	io.WriteString(res, "dog\n")
}

func me(res http.ResponseWriter, req *http.Request) {
	logRequest(req)
	io.WriteString(res, "Ives Ebneter\n")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":3000", nil)
}
