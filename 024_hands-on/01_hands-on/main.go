package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran\n")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("Error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "something.gohtml", "This is from dog")
	if err != nil {
		log.Fatalln("Error executing template", err)
	}
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "toby.jpg")
}
