package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var Upload = []byte(`
<html>
	<body>
	<form action="/upload" method="post"
	enctype="multipart/form-data">
		Image: <input type="file" name="my_file">
		<input type="submit" value="Upload">
	</form>
	</body>
</html>
`)

func mainpage(w http.ResponseWriter, r *http.Request) {
	w.Write(Upload)
}

func main8() {
	http.HandleFunc("/", mainpage)
	http.HandleFunc("/parse", handleparsgg)

	staticHandler := http.StripPrefix(
		"/data/",
		http.FileServer(http.Dir("./static")))

	http.Handle("/data/", staticHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handlepars(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * 1024 * 1025)

	file, handler, err := r.FormFile("my_file")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Fprintln(w, "Handler.filename", handler.Filename)
	fmt.Fprintln(w, "Handler.header", handler.Header)

	hash := md5.New()
	io.Copy(hash, file)
	fmt.Fprintf(w, "md5 %x\n", hash.Sum(nil))

}

/*
curl -v -X POST -H "Content-Type: application/json" -d '{"id": 2, "user": "ilya"}' http://localhost:8080/raw body
*/

type Params struct {
	ID   string
	User string
}

func handleparsgg(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	p := &Params{}

	err = json.Unmarshal(body, p)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "content-type %#v\n", r.Header.Get("Content-Type"))
	fmt.Fprintf(w, "params %#v\n", p)
}
