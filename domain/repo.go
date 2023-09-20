package domain

type Services interface {
	CreateRepoFromBaseTemplate(r Repo) (string, error)
	CreateServiceAndInfra(repoService Repo, repoInfra Repo) (*CreateServiceAndInfraResponse, error)
}

type Repo struct {
	Kind     string //kind of template being generated i.e.: golang-microservice
	RepoInfo RepoInfo
}

type RepoInfo struct {
	Owner              string `json:"owner"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	IncludeAllBranches bool   `json:"include_all_branches"`
	Private            bool   `json:"private"`
}
