package main

import "net/http"

func handler7(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
		Hello world! <br />
		<img src="/data/dofu.png" />
	`))
}

func main7() {
	http.HandleFunc("/", handler7)

	staticHandler := http.StripPrefix(
		"/data/",
		http.FileServer(http.Dir("./static")))

	http.Handle("/data/", staticHandler)

	http.ListenAndServe(":8080", nil)
}
