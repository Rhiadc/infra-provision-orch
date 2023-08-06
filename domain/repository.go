package domain

type Repo struct {
	Kind        string //kind of template that is being generated i.e.: golang-microservice
	Owner       string //default, company
	Name        string //name of the ms, will be preppendend with 'ms'
	Description string //description of the
}

type RepoInfo struct {
	Owner              string `json:"owner"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	IncludeAllBranches bool   `json:"include_all_branches"`
	Private            bool   `json:"private"`
}
