package domain

import (
	"fmt"

	"github.com/rhiadc/infra-provision-orch/config"
)

type Service struct {
	git config.Git
}

func NewService(git config.Git) *Service {
	return &Service{git: git}
}

func (s *Service) PushToGit(r Repo) {
	fmt.Println(r)
}
