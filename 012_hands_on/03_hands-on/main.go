package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	ZIP     int
	Region  string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		hotel{
			Name:    "Hotel Name",
			Address: "Hotel Address",
			City:    "City Name",
			Region:  "Region",
			ZIP:     123456,
		},
		hotel{
			Name:    "Hotel Name 2",
			Address: "Hotel Address 2",
			City:    "City Name 2",
			ZIP:     123456,
			Region:  "Region 2",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
