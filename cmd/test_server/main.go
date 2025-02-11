package main

import (
	"log"

	"github.com/mikerybka/api"
	"github.com/mikerybka/util"
)

type T struct {
	A string
	B int
	C bool
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
	s := &api.Server{
		Data: &T{},
	}
	log.Fatal(util.ListenAndServe(s))
}
