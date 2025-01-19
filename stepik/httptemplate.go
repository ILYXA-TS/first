package main

import (
	"net/http"
	"text/template"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

func main12() {
	tmp1 := template.Must(template.ParseFiles("users.html"))

	users := []User{
		{ID: 1, Name: "Ilya", Active: true},
		{ID: 2, Name: "<i>lexa<i>", Active: false},
		{ID: 3, Name: "vada", Active: true},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp1.Execute(w,
			struct {
				Users []User
			}{users})
	})
	http.ListenAndServe(":8080", nil)
}
