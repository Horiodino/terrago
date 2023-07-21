package web

import (
	"net/http"
	"text/template"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Person{Name: "John Doe", Age: 25, Country: "USA"}

		tmpl := template.Must(template.ParseFiles("template.html"))

		err := tmpl.Execute(w, person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
}
