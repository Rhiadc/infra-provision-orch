package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/pingcap/log"
	"github.com/rhiadc/infra-provision-orch/config"
)

const (
	GitURL     = "https://github.com"
	GitRepoAPI = "https://api.github.com/repos"
)

type Service struct {
	git config.Git
}

func NewService(git config.Git) *Service {
	return &Service{git: git}
}

func (s *Service) CreateRepoFromBaseTemplate(r Repo) (string, error) {
	url, token, err := s.getBaseRepo(r.Kind)
	if err != nil {
		return "", err
	}

	repoInfo := r.RepoInfo
	repoInfo.Private = false
	repoInfo.IncludeAllBranches = false
	repoInfo.Owner = s.git.Owner

	body, err := json.Marshal(repoInfo)
	if err != nil {
		return "", errors.New("Failed to unmarshall json")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", errors.New("Failed to create request")
	}

	// Set request headers.
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", s.git.APIVersion)

	// Perform the request.
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Error(fmt.Sprintf("Fail to perform request with error: %s", err.Error()))
		return "", errors.New("Failed to perform request")
	}

	defer resp.Body.Close()

	// Check the response status code.
	if resp.StatusCode != http.StatusCreated {
		log.Info(fmt.Sprintf("Github api status: %d", resp.StatusCode))
		return "", errors.New("Failed to create repository")
	}

	return fmt.Sprintf("%s/%s/%s", GitURL, s.git.Owner, r.RepoInfo.Name), nil
}

func (s Service) getBaseRepo(kind string) (string, string, error) {

	switch kind {
	case "ms":
		return s.git.GolangTemplateRepo, s.git.Token, nil
	case "infra":
		return s.git.InfraTemplateRepo, s.git.InfratemplateToken, nil
	}
	return "", "", ErrInvalidKindName
}

var ErrInvalidKindName = errors.New("Invalid kind name")
