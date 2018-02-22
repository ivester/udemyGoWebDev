package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type content struct {
	Text string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func logRequest(req *http.Request) {
	r := *req
	fmt.Println(r.Method, r.URL)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	logRequest(req)
	d := content{
		Text: "index",
	}
	err := tpl.ExecuteTemplate(res, "index.gohtml", d)
	handleError(res, err)
}

func dog(res http.ResponseWriter, req *http.Request) {
	logRequest(req)
	d := content{
		Text: "dog",
	}
	err := tpl.ExecuteTemplate(res, "dog.gohtml", d)
	handleError(res, err)
}

func me(res http.ResponseWriter, req *http.Request) {
	logRequest(req)
	d := content{
		Text: "Ives Ebneter",
	}
	err := tpl.ExecuteTemplate(res, "me.gohtml", d)
	handleError(res, err)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":3000", nil)
}
