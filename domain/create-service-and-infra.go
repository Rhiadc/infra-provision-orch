package domain

import (
	"fmt"

	"github.com/pingcap/log"
)

type CreateServiceAndInfraResponse struct {
	ResponseServiceReq string
	ResponseInfraReq   string
}

func (s *Service) CreateServiceAndInfra(repoService Repo, repoInfra Repo) (*CreateServiceAndInfraResponse, error) {
	res, err := s.CreateRepoFromBaseTemplate(repoService)
	if err != nil {
		return nil, err
	}
	reqRes := CreateServiceAndInfraResponse{ResponseServiceReq: res}

	res, err = s.CreateRepoFromBaseTemplate(repoInfra)
	if err != nil {
		log.Info("Error while creating the infra repo. Deleting the service repo - aborting...")
		if err := s.DeleteRepo(repoService.RepoInfo.Name); err != nil {
			log.Error(fmt.Sprintf("Error while deleting repo: %s ", err.Error()))
			return nil, err
		}
		return nil, err
	}
	reqRes.ResponseInfraReq = res
	return &reqRes, nil
}
