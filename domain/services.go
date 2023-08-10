package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rhiadc/infra-provision-orch/config"
)

type Service struct {
	git config.Git
}

func NewService(git config.Git) *Service {
	return &Service{git: git}
}

func (s *Service) PushToGit(r Repo) (string, error) {
	url, err := s.getBaseRepo(r.Kind)
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
	req.Header.Set("Authorization", "Bearer "+s.git.Token)
	req.Header.Set("X-GitHub-Api-Version", s.git.APIVersion)

	// Perform the request.
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("Failed to perform request")
	}
	defer resp.Body.Close()

	// Check the response status code.
	if resp.StatusCode != http.StatusCreated {
		return "", errors.New("Failed to create repository")
	}

	return repoInfo.Name, nil
}

func (s Service) getBaseRepo(kind string) (string, error) {

	switch kind {
	case "golang":
		return s.git.GolangTemplateRepo, nil
	}
	return "", errors.New("Invalid kind name")
}
