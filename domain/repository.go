package domain

import (
    "errors"
)

type Repository struct {
    Name string `json:"name"`
}

func CreateRepositoryFromMap(data map[string]interface{}) (Repository, error) {
    name, ok := data["name"].(string)
    if !ok {
        return Repository{}, errors.New("Can't read repository data, missing name.")
    }

    return Repository{Name: name}, nil
}
