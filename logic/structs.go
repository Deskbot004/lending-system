package logic

/*
Package contains relevant informations from the package

Name: Name of the package
Version: Latest installed version of the package
Require.Php: Version of php which is required for the package
Require.Typo3: Version of Typo3 which is required for the package
*/
type Package struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Requires struct {
		Php   string `json:"php"`
		Typo3 string `json:"typo3/cms-core"`
	} `json:"require"`
	Extra struct {
		Pack struct {
			Extkey string `json:"extension-key"`
		} `json:"typo3/cms"`
	} `json:"extra"`
}

/*
ProjectLock contains all packages of a project

Packages: Array of Package structs
ID: Project id from Gitlab as int
Name: (Optional) Name of the Project
*/
type ProjectLock struct {
	Packages []Package `json:"packages"`
	ID       int
	Name     string
}

/*
ProjectJSON extracts the important information from composer.json

Config: Structure containing the used php version
Require: Structure containing every required package
*/
type ProjectJSON struct {
	Config  map[string]map[string]string `json:"config"`
	Require map[string]string            `json:"require"`
}

/*
Stats contains the data given by packagist

Packages: Map of Versions of the package
Name: Name of the Package
*/
type Stats struct {
	Packages map[string][]Package `json:"packages"`
	Name     string
}

/*
Update saves the learnt information after compare

Own: Own version of a package
Latest: Latest version of a package
NeedUpdate: Boolean if the own package is outdated
PhpRequired: The Version of php that is required for the latest version
PhpHave: The php version the project has
CanUpdate: Boolean if the package can be upgraded to latest
*/
type Update struct {
	Own         string
	Latest      string
	NeedUpdate  bool
	PhpRequired string
	PhpHave     string
	CanUpdate   bool
}

/*
DBInfo is given as data for the database

Project: the relevant ProjectLock struct
Php: the Php version of the project
Typo3: the typo3 version of the project
Updates: array of all update data between project and packages
*/
type DBInfo struct {
	Project ProjectLock
	Php     string
	Typo3   string
	Updates map[string]Update
}

/*
BranchInfo containins information about a requested branch

Name: name of the branch
Default: boolean if the branch is the default branch
*/
type BranchInfo struct {
	Name    string `json:"name"`
	Default bool   `json:"default"`
}

/*
Commit containins information about a commit

Title: title of the commit
Author: author of the commit
Date: date of the commit
Link: Link to the TRON-ticket
*/
type Commit struct {
	Title  string `json:"title"`
	Author string `json:"author_name"`
	Date   string `json:"authored_date"`
	Link   string
}

/*
BranchCommit containins all commits of a branch

Commits: Array of all commits
*/
type BranchCommit struct {
	Commits []Commit `json:"commits"`
}

/*
Deployment is a singe deployment of a project

ProjectName: name of the project
ProjectID: id of the project
ID: id of the deployment
CreatedAt: Creation Date and Time of the Deployment
Environment: Contains Name and Link of the environment
Status: Contains the status of the deployment
*/
type Deployment struct {
	ProjectName string
	ProjectID   int
	ID          int         `json:"id"`
	CreatedAt   string      `json:"created_at"`
	Environment Environment `json:"environment"`
	Status      string      `json:"status"`
}

/*
Environment is a environment of a deployment

Name: name of the environment
Link: link to the environment
*/
type Environment struct {
	Name string `json:"name"`
	Link string `json:"external_url"`
}

/*
Deployments is an Array to read all Deployments from a request
*/
type Deployments []Deployment

/*
MergeRequests is an Array to read all MergeRequests from a request
*/
type MergeRequests []MergeRequest

/*
Author is a struct which saves the author data from a MergeRequest

Name: name of the author
Avatar: avatar of the author
Link: link to the gitlab of the author
*/
type Author struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar_url"`
	Link   string `json:"web_url"`
}

/*
MergeRequest is a single Merge request of an deployment

Name: name of the merge request
Author: author struct with further informations
CreatedAt: time and date of the request
*/
type MergeRequest struct {
	Name      string `json:"title"`
	Author    Author `json:"author"`
	CreatedAt string `json:"created_at"`
}
