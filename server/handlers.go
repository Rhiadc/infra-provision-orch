package server

import (
	"encoding/json"
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

// CreateRepo handler
// @Summary 	Create a repository for the type of service chosen
// @Description Returns a the link of the github project generated and the link of the infrastructure repository as well
// @Tags 		workflow
// @Accept		json
// @Produce		json
// @Param		request		body		createRepoRequest		true	"Create New repo"
// @Success		201			{object} 	createRepoResponse
// @Failure     400  		{object}    error
// @Router       /create-repo [post]
func (h Repohandler) CreateRepo(w http.ResponseWriter, r *http.Request) {
	data := &createRepoRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	//TODO: we need to pass the infra type dynamically, based on the service kind
	//it may be CLI infra, ms infra etc, so we need to identify the kind and create the
	//infra based on it
	infraRepoName := fmt.Sprintf("%s-infra", data.Name)
	dataInfra := &createRepoRequest{Kind: "infra", Name: infraRepoName, Description: infraRepoName}

	req := mapRepo(*data)
	reqDataInfra := mapRepo(*dataInfra)
	_, err := h.service.CreateServiceAndInfra(req, reqDataInfra)
	//res, err := h.service.CreateRepoFromBaseTemplate(req)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidKindName) {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	render.Status(r, http.StatusCreated)

	msg := createRepoResponse{fmt.Sprintf("Repo %s has been successfully created, with infra repo: %s", data.Name, dataInfra.Name)}
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type createRepoRequest struct {
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type createRepoResponse struct {
	message string
}
