package template

type ChiTemplates struct{}

func (c ChiTemplates) Main() []byte {
	return MainTemplate()
}

func (c ChiTemplates) Server() []byte {
	return MakeHTTPServer()
}

func (c ChiTemplates) Routes() []byte {
	return MakeChiRoutes()
}

func MakeChiRoutes() []byte {
	return []byte(`package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.helloWorldHandler)

	return r
}

func (s *Server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

`)
}
