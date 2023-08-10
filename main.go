package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/rhiadc/infra-provision-orch/config"
	"github.com/rhiadc/infra-provision-orch/domain"
	"github.com/rhiadc/infra-provision-orch/server"
)

func main() {
	r := chi.NewRouter()
	envs := config.LoadEnvVars()
	gitService := domain.NewService(envs.Git)
	server.NewRouter(r, gitService)

}
