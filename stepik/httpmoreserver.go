package main

import (
	"fmt"
	"net/http"
)

// 1
// func handle1(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("first"))
// }

// func main() {

// 	http.HandleFunc("/", handle1)

// 	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Second" + r.URL.String()))
// 	})
// 	http.HandleFunc("/pages/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("threee" + r.URL.String()))
// 	})

// 	fmt.Println("Server connect")
// 	http.ListenAndServe(":8080", nil)

// }

// 2

type Handler struct {
	Name string
}

func (s *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, s.Name, "Url:", r.URL.String())
}

func main3() {

	first := &Handler{Name: "first"}
	http.Handle("/test/", first)

	second := &Handler{Name: "Second"}
	http.Handle("/", second)

	fmt.Println("Starting server")
	http.ListenAndServe(":8080", nil)

}
