package api

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/mikerybka/util"
)

func isMap(v any) bool {
	kind := reflect.TypeOf(v).Kind()
	return kind == reflect.Map
}

type Map struct {
	Data map[string]any
}

func (s *Map) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	first, rest, _ := util.PopPath(r.URL.Path)

	if rest == "/" && r.Method == "PUT" {
		s.putItem(w, r)
		return
	}

	if rest == "/" && r.Method == "PATCH" {
		s.patchItem(w, r)
		return
	}

	if rest == "/" && r.Method == "DELETE" {
		s.deleteItem(r)
		return
	}

	next := &Server{
		Data: s.Data[first],
	}
	r.URL.Path = rest
	next.ServeHTTP(w, r)
}

func (s *Map) putItem(w http.ResponseWriter, r *http.Request) {
	first, _, _ := util.PopPath(r.URL.Path)
	err := json.NewDecoder(r.Body).Decode(s.Data[first])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (s *Map) patchItem(w http.ResponseWriter, r *http.Request) {
	first, _, _ := util.PopPath(r.URL.Path)
	err := json.NewDecoder(r.Body).Decode(s.Data[first])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (s *Map) deleteItem(r *http.Request) {
	first, _, _ := util.PopPath(r.URL.Path)
	delete(s.Data, first)
}
