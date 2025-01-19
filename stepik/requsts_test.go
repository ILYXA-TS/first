package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func main111() {

	http.HandleFunc("/", GetUser)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	key := r.FormValue("id")

	if key == "42" {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": 200, "resp: {"user":42}}`)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500, "err: "db_error"`)
	}

}

type Testing struct {
	ID         string
	Response   string
	StatusCode int
}

func TestGetUSer(t *testing.T) {
	cases := []Testing{
		Testing{
			ID:         "42",
			Response:   `{"status": 200, "resp: {"user":42}}`,
			StatusCode: http.StatusOK},
		Testing{
			ID:         "500",
			Response:   `{"status": 500, "err: "db_error"`,
			StatusCode: http.StatusInternalServerError},
	}

	for chais, cs := range cases {
		url := "http://example.com/api/user?id=" + cs.ID
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()

		GetUser(w, req)

		if w.Code != cs.StatusCode {
			t.Errorf("gg%d", chais)

		}
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		bodystring := string(body)
		if bodystring != cs.Response {
			t.Errorf("gg2%d", chais)
		}
	}
}
