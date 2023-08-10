package domain

type Services interface {
	PushToGit(r Repo) (string, error)
}

type Repo struct {
	Kind     string //kind of template being generated i.e.: golang-microservice
	RepoInfo RepoInfo
}

type RepoInfo struct {
	Owner              string //default, company
	Name               string //ms name, will be preppendend with 'ms'
	Description        string
	IncludeAllBranches bool
	Private            bool
}
