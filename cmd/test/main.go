package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/mikerybka/api"
)

type Test struct {
	Method       string
	Path         string
	Body         string
	ExpectedCode int
	ExpectedBody string
}

type T struct {
	A string
	B int
	C bool
	D map[string]string
	E map[string]bool
	F map[string]int
	G map[string]S
	H *S
	I []string
	J []S
	K []int
}

type S struct {
	A int
	B map[string]struct {
		D string
		E int
		F bool
	}
}

func (t *T) SetA(s string) {
	t.A = s
}

func (t *T) Inc() {
	t.B += 1
}

func (t *T) Dec() {
	t.B -= 1
}

func (t *T) Switch() {
	t.C = !t.C
}

func main() {
	// go func() {
	s := &api.Server{
		Data: &T{},
	}
	panic(http.ListenAndServe(":2999", s))
	// }()

	// tests := []Test{
	// 	// Smoke test
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/",
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `{"A":"","B":0,"C":false}` + "\n",
	// 	},

	// 	// Test getting struct fields
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/a",
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `""` + "\n",
	// 	},
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/A",
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `""` + "\n",
	// 	},
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/b",
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `0` + "\n",
	// 	},
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/B",
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `0` + "\n",
	// 	},
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/c",
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `false` + "\n",
	// 	},
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/C",
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `false` + "\n",
	// 	},
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/d",
	// 		ExpectedCode: 404,
	// 		ExpectedBody: "404 page not found" + "\n",
	// 	},

	// 	// Test method calls
	// 	{
	// 		Method:       "POST",
	// 		Path:         "/set-a",
	// 		Body:         `["\"yellow\""]`,
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `[]` + "\n",
	// 	},
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/",
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `{"A":"yellow","B":0,"C":false}` + "\n",
	// 	},
	// 	{
	// 		Method:       "POST",
	// 		Path:         "/inc",
	// 		Body:         `[]`,
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `[]` + "\n",
	// 	},
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/",
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `{"A":"yellow","B":1,"C":false}` + "\n",
	// 	},
	// 	{
	// 		Method:       "POST",
	// 		Path:         "/switch",
	// 		Body:         `[]`,
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `[]` + "\n",
	// 	},
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/",
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `{"A":"yellow","B":1,"C":true}` + "\n",
	// 	},
	// 	{
	// 		Method:       "POST",
	// 		Path:         "/dec",
	// 		Body:         `[]`,
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `[]` + "\n",
	// 	},
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/",
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `{"A":"yellow","B":0,"C":true}` + "\n",
	// 	},
	// 	{
	// 		Method:       "POST",
	// 		Path:         "/switch",
	// 		Body:         `[]`,
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `[]` + "\n",
	// 	},
	// 	{
	// 		Method:       "GET",
	// 		Path:         "/",
	// 		ExpectedCode: 200,
	// 		ExpectedBody: `{"A":"yellow","B":0,"C":false}` + "\n",
	// 	},

	// 	// TODO: Test custom ServeHTTP
	// 	// TODO: Test PATCH/PUT
	// 	// TODO: Test arrays
	// 	// TODO: Test maps

	// 	// Test depth
	// }
	// for i, t := range tests {
	// 	code, body := send(t.Method, "http://localhost:2999"+t.Path, t.Body)
	// 	if code != t.ExpectedCode {
	// 		fmt.Println("ERROR:", "test", i, "expected code", t.ExpectedCode, "got code", code)
	// 		fmt.Println("Body:", strings.TrimSpace(body))
	// 		return
	// 	}
	// 	if body != t.ExpectedBody {
	// 		fmt.Println("ERROR:", "test", i, "expected body", strings.TrimSpace(t.ExpectedBody), "got body", strings.TrimSpace(body))
	// 		return
	// 	}
	// }
	// os.Exit(0)
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
