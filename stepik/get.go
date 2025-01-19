package main

import (
	"fmt"
	"net/http"
)

// get parametri
// http://127.0.0.1:8080/?param=1&key=2

func handle4(w http.ResponseWriter, r *http.Request) {
	myparam := r.URL.Query().Get("param")
	if myparam != "" {
		fmt.Fprintln(w, "myparam is ", myparam)
	}

	myparam2 := r.FormValue("key")
	if myparam2 != "" {
		fmt.Fprintln(w, "my param2", myparam2)
	}

}

func main5() {
	http.HandleFunc("/", postparam)

	http.ListenAndServe(":8080", nil)
}

// post parametri

var login = []byte(`
<html>
	<body>
	<form action="/" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
	</body>
</html>	
`)

func postparam(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Write(login)
		return
	}

	// r.ParseForm()
	// inputlogin := r.Form["login"][0]

	inputlogin := r.FormValue("login")

	fmt.Fprintln(w, "your enter", inputlogin)

}
