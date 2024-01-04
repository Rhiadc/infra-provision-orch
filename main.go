package main

import (
	"github.com/rhiadc/infra-provision-orch/config"
	"github.com/rhiadc/infra-provision-orch/domain"
	"github.com/rhiadc/infra-provision-orch/server"
)

// @title Infra provision orchestrator API
// @version 1.0
// @description This is the API for responsible for generating a microservice repository (based on the language chosen) with the infrastrucure repository as well
// @termsOfService http://swagger.io/terms/

// @contact.name rhiadc
// @contact.email rhiad.ciccoli@gmail.com

// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {

	envs := config.LoadEnvVars()
	gitService := domain.NewService(envs.Git)
	_ = server.NewServer(gitService, envs)

}
