package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Cart struct {
	Pay string
}

type testcase struct {
	ID     string
	Result *Chekoutresult
	err    bool
}
type Chekoutresult struct {
	status  int
	balance int
	err     string
}

func (s *Cart) Chekout(id string) (*Chekoutresult, error) {
	url := s.Pay + "?id=" + id

	resp, err := http.Get(url)
	if err != nil {
		return &Chekoutresult{}, err
	}

	defer resp.Body.Close()
	es, _ := ioutil.ReadAll(resp.Body)

	gg := &Chekoutresult{}

	json.Unmarshal(es, gg)

	return gg, nil

}

func damm(w http.ResponseWriter, r *http.Request) {

	es := r.FormValue("id")

	switch es {
	case "42":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": 200, "balance": 100500}`)
	case "100500":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": 400, "err": bad_balance}`)
	case "__broken_json":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": 400}`)
	case "__internal_error":
		fallthrough
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Test(t *testing.T) {

	cases := []testcase{
		{
			ID: "42",
			Result: &Chekoutresult{
				status:  200,
				balance: 100500,
				err:     "",
			},
			err: false},
		{
			ID: "100500",
			Result: &Chekoutresult{
				status:  400,
				balance: 0,
				err:     "bad_balance",
			},
			err: false},
		{
			ID:     "__broken_json",
			Result: nil,
			err:    true},

		{ID: "__internal_error",
			Result: nil,
			err:    true},
	}
	ts := httptest.NewServer(http.HandlerFunc(damm))

	for casenam, item := range cases {
		c := &Cart{
			Pay: ts.URL,
		}

		result, err := c.Chekout(item.ID)

		if err != nil && !item.err {
			t.Error("gg3")
		}

		if err == nil && !item.err {
			t.Error("gg3")
		}
		if !reflect.DeepEqual(item.Result, result) {
			t.Error("gg3", casenam)
		}
	}
	ts.Close()
}
