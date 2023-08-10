package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pingcap/log"
	"github.com/rhiadc/infra-provision-orch/domain"
)

type Repohandler struct {
	service domain.Services
}

func NewRepoHandler(service domain.Services) *Repohandler {
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
	res, err := h.service.PushToGit(req)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	render.Status(r, http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("%s has been successfully created", res)))
}

type createRepoRequest struct {
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type createRepoResponse struct {
}
