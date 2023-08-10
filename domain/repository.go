package domain

type Repo struct {
	Kind               string //kind of template being generated i.e.: golang-microservice
	Owner              string //default, company
	Name               string //ms name, will be preppendend with 'ms'
	Description        string
	IncludeAllBranches bool
	Private            bool
}
