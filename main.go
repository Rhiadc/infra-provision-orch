package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/rhiadc/infra-provision-orch/domain"
	"github.com/rhiadc/infra-provision-orch/server"
)

func main() {
	r := chi.NewRouter()
	gitService := domain.NewService()
	server.NewRouter(r, gitService)

}
