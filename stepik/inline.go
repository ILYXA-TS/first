package main

import (
	"net/http"
	"text/template"
)

type tplparam struct {
	URL     string
	Browser string
}

const EXAMPL = `
Browser {{.Browser}}

you at {{.URL}}
`

func main123() {
	http.HandleFunc("/", handlelin)

	http.ListenAndServe(":8080", nil)
}

func handlelin(w http.ResponseWriter, r *http.Request) {
	tmp1 := template.New(`exampl`)
	tmp1, _ = tmp1.Parse(EXAMPL)

	parms := &tplparam{
		URL:     r.URL.String(),
		Browser: r.UserAgent(),
	}
	tmp1.Execute(w, parms)
}
