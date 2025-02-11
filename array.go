package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/mikerybka/util"
)

func isArray(v any) bool {
	kind := reflect.TypeOf(v).Kind()
	return kind == reflect.Array || kind == reflect.Slice
}

type Array struct {
	Data []any
}

func (s *Array) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	first, rest, isRoot := util.PopPath(r.URL.Path)

	if isRoot && r.Method == "POST" {
		s.append(w, r)
		return
	}

	if rest == "/" && r.Method == "DELETE" {
		s.delete(w, r)
		return
	}

	i, _ := strconv.Atoi(first)
	if i < 0 || i >= len(s.Data) {
		http.NotFound(w, r)
		return
	}

	next := &Server{
		Data: s.Data[i],
	}
	r.URL.Path = rest
	next.ServeHTTP(w, r)
}

func (s *Array) append(w http.ResponseWriter, r *http.Request) {
	t := reflect.TypeOf(s.Data).Elem()
	v := reflect.New(t)
	json.NewDecoder(r.Body).Decode(v.Interface())
	s.Data = append(s.Data, v.Interface())
	i := len(s.Data) - 1
	fmt.Fprintf(w, "%d\n", i)
}

func (s *Array) delete(w http.ResponseWriter, r *http.Request) {
	first, _, _ := util.PopPath(r.URL.Path)
	i, err := strconv.Atoi(first)
	if err != nil || i < 0 || i >= len(s.Data) {
		http.NotFound(w, r)
		return
	}
	s.Data = append(s.Data[:i], s.Data[i+1:]...)
	fmt.Fprintln(w, len(s.Data))
}
