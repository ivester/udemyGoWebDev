package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "fileOne.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "fileTwo.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
