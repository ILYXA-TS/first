package main

import (
	"fmt"
	"net/http"
)

func main2() {
	http.HandleFunc("/", handle)

	fmt.Println("Starting server")

	http.ListenAndServe(":8080", nil)

}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello user")
	w.Write([]byte("hello"))
}
