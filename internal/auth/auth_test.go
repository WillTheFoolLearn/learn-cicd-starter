package auth

import (
	"net/http"
	"testing"
)

func Test1(t *testing.T) {
	header := make(http.Header)
	header1 := make(http.Header)
	header2 := make(http.Header)

	header1.Add("Authorization", "Bearer cheese")
	header2.Add("Cheese", "Authorization test")

	testList := []Tests{
		{
			Name:   "Blank header",
			Header: header,
			Result: false,
		},
		{
			Name:   "Valid header",
			Header: header1,
			Result: true,
		},
		{
			Name:   "Wrong header",
			Header: header2,
			Result: false,
		},
	}

	for _, test := range testList {
		testResult, err := GetAPIKey(test.Header)
		if !(err != nil) != test.Result || !(testResult == "") != test.Result {
			t.Errorf("Test %s failed: %v", test.Name, err)
		}
	}
}

type Tests struct {
	Name   string
	Header http.Header
	Result bool
}
