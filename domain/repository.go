package domain

import (
    "errors"
)

//  A repository represents the "naked" repository data without any further
//  repository handling like clone, commit or push.
type Repository struct {
    Name string `json:"name"`
    CloneURL string `json:"clone_url"`
    Owner string `json:"owner"`
}

func CreateRepositoryFromMap(data map[string]interface{}) (Repository, error) {
    name, ok := data["name"].(string)
    if !ok {
        return Repository{}, errors.New("Can't read repository data, missing name.")
    }

    cloneUrl, ok := data["clone_url"].(string)
    if !ok {
        return Repository{}, errors.New("Can't read repository data, missing clone url.")
    }

    owner, ok := data["owner"].(map[string]interface{})
    if !ok {
        return Repository{}, errors.New("Can't read repository data, missing owner information.")
    }

    ownerName, ok := owner["name"].(string)
    if !ok {
        ownerName, ok = owner["login"].(string)
        if !ok {
            return Repository{}, errors.New("Can't read repository data, missing owner name.")
        }
    }

    return Repository{Name: name, CloneURL: cloneUrl, Owner: ownerName}, nil
}
