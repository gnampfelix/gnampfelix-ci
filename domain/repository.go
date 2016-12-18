package domain

import (
    "errors"
)

type Repository struct {
    Name string `json:"name"`
    CloneURL string `json:"clone_url"`
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

    return Repository{Name: name, CloneURL: cloneUrl}, nil
}
