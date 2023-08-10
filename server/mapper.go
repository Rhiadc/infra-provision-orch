package server

import "github.com/rhiadc/infra-provision-orch/domain"

func mapRepo(c createRepoRequest) domain.Repo {
	return domain.Repo{Kind: c.Kind, RepoInfo: domain.RepoInfo{Name: c.Name, Description: c.Description}}
}
