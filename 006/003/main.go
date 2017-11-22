package main

import "html/template"
import "os"
import "log"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("someFile.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "someFile.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
