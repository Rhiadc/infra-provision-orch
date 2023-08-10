package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rhiadc/infra-provision-orch/domain"
)

type Repohandler struct {
	service *domain.Service
}

func NewRepoHandler(service *domain.Service) *Repohandler {
	return &Repohandler{service: service}
}

func (c *createRepoRequest) Bind(r *http.Request) error {

	if c.Kind == "" {
		return errors.New("missing required Kind field.")
	}

	if c.Name == "" {
		return errors.New("missing required Name field.")
	}

	return nil
}

func (h Repohandler) CreateRepo(w http.ResponseWriter, r *http.Request) {
	data := &createRepoRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	req := mapRepo(*data)
	h.service.PushToGit(req)
	render.Status(r, http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("%s has been created successfully", data.Name)))
}

type createRepoRequest struct {
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type createRepoResponse struct {
}
