package logic

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

/*
ConvertToStruct converts the given json file into the a given struct construct
which can be more easily accessed

response: json from the Get request
f: Struct genric to which the json gets converted
*/
func ConvertToStruct[T any](response *http.Response, f *T) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal([]byte(body), f)
	if err != nil {
		if !strings.Contains(err.Error(), "cannot unmarshal") && !strings.Contains(err.Error(), "into Go") {
			log.Println(err)
		}
	}
}

/*
RemoveUnnecessary removes packages found in lock which are not found in json

projectLock: Project struct from .lock
projectJson: Project struct from .json
*/
func RemoveUnnecessary(projectLock *ProjectLock, projectJson *ProjectJSON) {
	var newList []Package

	for k := range projectJson.Require {
		removeUnnecessaryHelp(projectLock, projectJson, &newList, k)
	}
	projectLock.Packages = newList
}

func removeUnnecessaryHelp(projectLock *ProjectLock, projectJson *ProjectJSON, newList *[]Package, k string) {
	// remove all ext/ cause Benni said so
	if strings.Contains(k, "ext") {
		delete(projectJson.Require, k)
		return
	}

	// iterate over all Projects and only keep those found in .json
	for i := 0; i < len(projectLock.Packages); i++ {
		if projectLock.Packages[i].Name == k {
			*newList = append(*newList, projectLock.Packages[i])
			break
		}
	}
}

/*
FindMainBranch searches for the main branch by searching for default

search: Array containing the relevant branch information as struct
return: if successful the name of the main branch, "" else
*/
func FindMainBranch(search []BranchInfo) string {
	for i := 0; i < len(search); i++ {
		if search[i].Default {
			return search[i].Name
		}
	}
	return ""
}
