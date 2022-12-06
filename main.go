package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Easy", Done: false},
				{Title: "Medium", Done: true},
				{Title: "Hard", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPreFix("/assests/", static))

}
