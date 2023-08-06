package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rhiadc/infra-provision-orch/domain"
)

func NewRouter(r *chi.Mux, service *domain.Service) {
	repoHandler := NewRepoHandler(service)
	r.Post("/{param}", repoHandler.CreateRepo)
	//go func() {
	if err := http.ListenAndServe(":3333", r); err != nil {
		log.Fatal(err)
	}
	//}()
}
