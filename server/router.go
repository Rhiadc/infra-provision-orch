package server

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/rhiadc/infra-provision-orch/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (s *Server) router(r *chi.Mux) {
	repoHandler := NewRepoHandler(s.service)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))
	r.Post("/create-repo", repoHandler.CreateRepo)
}
