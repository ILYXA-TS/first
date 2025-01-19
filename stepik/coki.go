package main

import (
	"fmt"
	"net/http"
	"time"
)

func main6() {

	// http.HandleFunc("/login", handlelogin)
	// http.HandleFunc("/logout", handlelogout)
	// http.HandleFunc("/", handlemain)

	http.HandleFunc("/", handle5)

	http.ListenAndServe(":8080", nil)
}

func handlemain(w http.ResponseWriter, r *http.Request) {

	seson, err := r.Cookie("session_id")

	login := (err != http.ErrNoCookie)

	if login {
		fmt.Fprintln(w, `<a href="logout">logout</a>`)
		fmt.Fprintln(w, "Welcom, "+seson.Value)
	} else {
		fmt.Fprintln(w, `<a href="login">login</a>`)
		fmt.Fprintln(w, "You need to login")
	}
}

func handlelogin(w http.ResponseWriter, r *http.Request) {
	exprethion := time.Now().Add(10 * time.Hour)

	cokie := http.Cookie{
		Name:    "session_id",
		Value:   "ilya",
		Expires: exprethion,
	}
	http.SetCookie(w, &cokie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func handlelogout(w http.ResponseWriter, r *http.Request) {

	sesion, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	sesion.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, sesion)
	http.Redirect(w, r, "/", http.StatusFound)

}

func handle5(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ReqestID", "rtkhjnidfhidfug4")

	fmt.Fprintln(w, "You broswer is", r.UserAgent())
	fmt.Fprintln(w, "Your accept is", r.Header.Get("Accept"))
}
