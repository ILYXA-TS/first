package main

import (
	"net/http"
	"text/template"
)

type User1 struct {
	ID     int
	Name   string
	Active bool
}

func main() {
	tmp1, _ := template.New("").ParseFiles("users2.html")

	users := []User1{
		{ID: 1, Name: "Ilya", Active: true},
		{ID: 2, Name: "lexa", Active: false},
		{ID: 3, Name: "vada", Active: true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp1.ExecuteTemplate(w, "users2.html",
			struct {
				Users []User1
			}{users})
	})
	http.ListenAndServe(":8080", nil)
}

func (s *User1) Printactive() string {
	if !s.Active {
		return ""
	}
	return "method" + s.Name + " active"
}
