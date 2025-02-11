package api

import (
	"net/http"
	"reflect"

	"github.com/mikerybka/util"
)

func isStruct(v any) bool {
	kind := reflect.TypeOf(v).Kind()
	return kind == reflect.Struct
}

type Struct struct {
	Data any
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	first, rest, _ := util.PopPath(r.URL.Path)

	// Check each field
	t := reflect.TypeOf(s.Data)
	for {
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		} else {
			break
		}
	}
	v := reflect.ValueOf(s.Data)
	for {
		if v.Kind() == reflect.Pointer {
			v = v.Elem()
		} else {
			break
		}
	}
	pascalCaseFieldName := kebab2pascal(first)
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		if pascalCaseFieldName == fieldName {
			next := &Server{
				Data: v.Field(i).Interface(),
			}
			r.URL.Path = rest
			next.ServeHTTP(w, r)
			return
		}
	}

	http.NotFound(w, r)
}
