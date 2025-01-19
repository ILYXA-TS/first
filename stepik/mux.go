package main

import (
	"fmt"
	"net/http"
)

// 1
// func handler3(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(r.URL.String()))
// }

// func main() {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/", handler3)

// 	server := http.Server{
// 		Addr:         ":8080",
// 		Handler:      mux,
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 	}

// 	server.ListenAndServe()
// }

//2

func runserver(server string) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Addr:", server, "URL:", r.URL.String())
	})

	server1 := http.Server{
		Addr:    server,
		Handler: mux,
	}

	fmt.Println("starting server")
	server1.ListenAndServe()

}

func main4() {
	go runserver(":8081")
	runserver(":8080")
}
