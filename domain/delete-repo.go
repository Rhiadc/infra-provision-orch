package domain

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/pingcap/log"
)

func (s *Service) DeleteRepo(repo string) error {
	url := fmt.Sprintf("%s/%s/%s", GitRepoAPI, s.git.Owner, repo)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return errors.New("Failed to create request")
	}

	// Set request headers.
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+s.git.GlobalToken)
	req.Header.Set("X-GitHub-Api-Version", s.git.APIVersion)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusNoContent {
		errMsg := resp.Status
		if err != nil {
			errMsg = err.Error()
		}
		log.Error(fmt.Sprintf("Fail to perform delete request with error: %s", errMsg))
		return errors.New("Failed to perform request")
	}
	defer resp.Body.Close()

	return nil
}
