package api

import "net/http"

type Server[RootType any] struct {
	Data *RootType
}

func (s *Server[RootType]) ServeHTTP(w http.ResponseWriter, r *http.Request)
