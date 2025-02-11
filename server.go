package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/mikerybka/util"
)

type Server struct {
	Data any
}

func (s *Server) callMethod(w http.ResponseWriter, r *http.Request) {
	// Get the method to call
	first, _, _ := util.PopPath(r.URL.Path)
	t := reflect.TypeOf(s.Data)
	mName := kebab2pascal(first)
	m, ok := t.MethodByName(mName)
	if !ok {
		panic("no method " + mName)
	}

	// Read raw inputs as a string list
	rawArgs := []string{}
	err := json.NewDecoder(r.Body).Decode(&rawArgs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert string args to reflect.Values
	inLen := m.Type.NumIn()
	args := make([]reflect.Value, inLen)
	args[0] = reflect.ValueOf(s.Data)
	for i := 1; i < inLen; i++ {
		inType := m.Type.In(i)
		val := reflect.New(inType)
		rawJSON := []byte(rawArgs[i-1])
		err := json.Unmarshal(rawJSON, val.Interface())
		if err != nil {
			http.Error(w, fmt.Sprintf("arg %d: %s", i, err), http.StatusBadRequest)
			return
		}
		args[i] = val.Elem()
	}

	// Call method
	results := m.Func.Call(args)

	// Convert results to string list
	res := []string{}
	for _, re := range results {
		b, err := json.Marshal(re.Interface())
		if err != nil {
			panic(err)
		}
		res = append(res, string(b))
	}

	// Write JSON output
	json.NewEncoder(w).Encode(res)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Use ServeHTTP method if type is an http.Handler
	if hasServeHTTP(s.Data) {
		serveHTTP(s.Data, w, r)
		return
	}

	// Handle method calls
	first, rest, isRoot := util.PopPath(r.URL.Path)
	if rest == "/" && r.Method == "POST" && hasMethod(s.Data, kebab2pascal(first)) {
		s.callMethod(w, r)
		return
	}

	// Handle GET / requests
	if isRoot && r.Method == "GET" {
		s.getRoot(w)
		return
	}

	// Handle PATCH / requests
	if isRoot && r.Method == "PATCH" {
		s.patchRoot(w, r)
		return
	}

	// Handle PUT / requests
	if isRoot && r.Method == "PUT" {
		s.patchRoot(w, r) // TODO: implement putRoot
		return
	}

	// Split logic based on type kind
	t := reflect.TypeOf(s.Data)
	for {
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		} else {
			break
		}
	}
	kind := t.Kind()

	if kind == reflect.Array || kind == reflect.Slice {
		h := &Array{
			Data: s.Data.([]any),
		}
		h.ServeHTTP(w, r)
		return
	}

	if kind == reflect.Map {
		h := &Map{
			Data: s.Data.(map[string]any),
		}
		h.ServeHTTP(w, r)
		return
	}

	if kind == reflect.Struct {
		h := &Struct{
			Data: s.Data,
		}
		h.ServeHTTP(w, r)
		return
	}

	http.NotFound(w, r)
}

func (s *Server) getRoot(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(s.Data)
}

func (s *Server) patchRoot(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(s.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func hasMethod(v any, name string) bool {
	t := reflect.TypeOf(v)
	_, ok := t.MethodByName(name)
	return ok
}

// returns true if v is an http.Handler
func hasServeHTTP(v any) bool {
	t := reflect.TypeOf(v)
	m, ok := t.MethodByName("ServeHTTP")
	if !ok {
		return false
	}
	if m.Type.NumIn() != 3 {
		return false
	}
	if m.Type.In(1).String() != "http.ResponseWriter" {
		return false
	}
	if m.Type.In(2).String() != "*http.Request" {
		return false
	}
	return true
}

func serveHTTP(v any, w http.ResponseWriter, r *http.Request) {
	v.(http.Handler).ServeHTTP(w, r)
}
