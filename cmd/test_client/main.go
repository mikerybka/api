package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mikerybka/util"
)

type Test struct {
	Method       string
	Path         string
	Body         string
	ExpectedCode int
	ExpectedBody string
}

func main() {
	backendURL := util.RequireEnvVar("BACKEND_URL")
	tests := []Test{
		// Smoke test
		{
			Method:       "GET",
			Path:         "/",
			ExpectedCode: 200,
			ExpectedBody: `{"A":"","B":0,"C":false}` + "\n",
		},
		{
			Method:       "GET",
			Path:         "/a",
			ExpectedCode: 200,
			ExpectedBody: `""` + "\n",
		},
		{
			Method:       "GET",
			Path:         "/A",
			ExpectedCode: 200,
			ExpectedBody: `""` + "\n",
		},
		{
			Method:       "GET",
			Path:         "/b",
			ExpectedCode: 200,
			ExpectedBody: `0` + "\n",
		},
		{
			Method:       "GET",
			Path:         "/B",
			ExpectedCode: 200,
			ExpectedBody: `0` + "\n",
		},
		{
			Method:       "GET",
			Path:         "/c",
			ExpectedCode: 200,
			ExpectedBody: `false` + "\n",
		},
		{
			Method:       "GET",
			Path:         "/C",
			ExpectedCode: 200,
			ExpectedBody: `false` + "\n",
		},
		{
			Method:       "GET",
			Path:         "/d",
			ExpectedCode: 404,
			ExpectedBody: "404 page not found" + "\n",
		},

		// Test method calls
		{
			Method:       "POST",
			Path:         "/set-a",
			Body:         `["\"yellow\""]`,
			ExpectedCode: 200,
			ExpectedBody: `[]` + "\n",
		},
	}
	for i, t := range tests {
		code, body := send(t.Method, backendURL+t.Path, t.Body)
		if code != t.ExpectedCode {
			fmt.Println("ERROR:", "test", i, "expected code", t.ExpectedCode, "got code", code)
			fmt.Println("Body:", strings.TrimSpace(body))
			return
		}
		if body != t.ExpectedBody {
			fmt.Println("ERROR:", "test", i, "expected body", strings.TrimSpace(t.ExpectedBody), "got body", strings.TrimSpace(body))
			return
		}
	}
}

func send(method, url, body string) (int, string) {
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	b, _ := io.ReadAll(res.Body)
	return res.StatusCode, string(b)
}
