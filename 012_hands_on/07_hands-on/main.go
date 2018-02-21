package main

import "html/template"
import "os"

var tpl *template.Template

type menu struct {
	Starter string
	Main    string
}

type dayTime struct {
	Name string
	Menu []menu
}

type restaurant struct {
	Breakfast dayTime
	Lunch     dayTime
	Dinner    dayTime
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := []restaurant{
		restaurant{
			Breakfast: dayTime{
				Name: "Breakfast",
				Menu: []menu{
					menu{
						Starter: "Starter Breakfast 1",
						Main:    "Main Breakfast 1",
					},
					menu{
						Starter: "Starter Breakfast 2",
						Main:    "Main Breakfast 2",
					},
				},
			},
			Lunch: dayTime{
				Name: "Lunch",
				Menu: []menu{
					menu{
						Starter: "Starter Lunch 1",
						Main:    "Main Lunch 1",
					},
					menu{
						Starter: "Starter Lunch 2",
						Main:    "Main Lunch 2",
					},
				},
			},
			Dinner: dayTime{
				Name: "Dinner",
				Menu: []menu{
					menu{
						Starter: "Starter Dinner 1",
						Main:    "Main Dinner 1",
					},
					menu{
						Starter: "Starter Dinner 2",
						Main:    "Main Dinner 2",
					},
				},
			},
		},
		restaurant{
			Breakfast: dayTime{
				Name: "2 Breakfast",
				Menu: []menu{
					menu{
						Starter: "Starter Breakfast 1",
						Main:    "Main Breakfast 1",
					},
					menu{
						Starter: "Starter Breakfast 2",
						Main:    "Main Breakfast 2",
					},
				},
			},
			Lunch: dayTime{
				Name: "2 Lunch",
				Menu: []menu{
					menu{
						Starter: "Starter Lunch 1",
						Main:    "Main Lunch 1",
					},
					menu{
						Starter: "Starter Lunch 2",
						Main:    "Main Lunch 2",
					},
				},
			},
			Dinner: dayTime{
				Name: "2 Dinner",
				Menu: []menu{
					menu{
						Starter: "Starter Dinner 1",
						Main:    "Main Dinner 1",
					},
					menu{
						Starter: "Starter Dinner 2",
						Main:    "Main Dinner 2",
					},
				},
			},
		},
	}

	tpl.Execute(os.Stdout, r)
}
