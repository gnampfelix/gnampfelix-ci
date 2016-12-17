package events

import (
    "github.com/Thamtham/gnampfelix-ci/domain"
    "github.com/Thamtham/gnampfelix-ci/config"
    "encoding/json"
    "errors"
    "fmt"
)

//  GitHub sends a push notification after each push on the repo the webhook is
//  assigned to. The type Push represents this notification for further computation.
type Push struct {
    Ref string
    Repository domain.Repository
}

//  Push implements Unmarshaler. This is due to the fact that golang currently
//  does not support "required" fields for Unmarshal. A push needs a repository
//  and a ref!
func (p *Push)UnmarshalJSON(data []byte) error {
    var resultMap map[string]interface{}
    err := json.Unmarshal(data, &resultMap)
    if err != nil {
        return err
    }

    ref, ok := resultMap["ref"].(string)
    if !ok {
        return errors.New("The push could not be read, missing ref.")
    }

    repoMap, ok := resultMap["repository"].(map[string]interface{})
    if !ok {
        return errors.New("The push could not be read, missing repository information.")
    }

    repository, err := domain.CreateRepositoryFromMap(repoMap)
    if err != nil {
        return err
    }

    p.Ref = ref
    p.Repository = repository
    return nil
}

func (p Push)HandleEvent() error {
    configFile, err := config.ReadRepoConfig(p.Repository.Name)
    if err != nil {
        return err
    }
    action, err := configFile.GetAction(p.Ref, "push")
    if err != nil {
        return err
    }
    tmp, err := action.Run()
    fmt.Println(tmp)
    return err
}
