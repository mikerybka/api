package api

import "github.com/mikerybka/golang"

func GenerateGoServer(t, path string) string {
	f := &golang.File{
		PkgName: "main",
	}
	// f.AddFunc("main", golang.Func{
	// 	Body: []golang.Stmt{
	// 		{
	// 			IsDeclStmt: true,
	// 			DeclStmt:   &golang.DeclStmt{},
	// 		},
	// 	},
	// })
	return f.String()
}

var goServerTmpl string = `package main

import (
	"log"

	"github.com/mikerybka/api"
	"github.com/mikerybka/util"

	"%s"
)

func main() {
	s := &api.Server[%s]{}
	log.Fatal(util.ListenAndServe(s))
}
`
