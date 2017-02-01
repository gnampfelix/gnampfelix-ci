package events

import (
    "encoding/json"
    "github.com/gnampfelix/gnampfelix-ci/domain"
    "github.com/gnampfelix/gnampfelix-ci/config"
    "errors"
)

//  A Ping is sent by GitHub after a webhook is created. The type Ping contains
//  all data that is needed for verification.
type Ping struct {
    Repository domain.Repository `json:"repository"`
    Hook domain.Hook `json:"hook"`
}

//  Ping implements Unmarshaler. This is due to the fact that golang currently
//  does not support "required" fields for Unmarshal. A ping needs a repository
//  and a hook!
func (p *Ping)UnmarshalJSON(data []byte) error {
    var result map[string]interface{}
    var ok bool

    err := json.Unmarshal(data, &result)
    if err != nil {
        return err
    }

    repositoryData, ok := result["repository"].(map[string]interface{})
    if !ok {
        return errors.New("Can't read ping event, invalid repository data.")
    }

    hookData, ok := result["hook"].(map[string]interface{})
    if !ok {
        return errors.New("Can't read ping event, invalid hook data.")
    }

    repo, err := domain.CreateRepositoryFromMap(repositoryData)
    if err != nil {
        return err
    }

    hook, err := domain.CreateHookFromMap(hookData)
    if err != nil {
        return err
    }

    p.Repository = repo
    p.Hook = hook

    return nil
}

//  HasValidConfiguration() verifies if the user has already set up a valid
//  configuration that meets the data sent via the ping.
//  A valid configuration is defined by the following:
//      - "{repoName}.json" in the ci rootFolder
//      - valid handling for each event that is noted in the ping
//  A valid handling is determined by the config file itself.
func (p Ping)HasValidConfiguration() bool {
    repoConfig, err := config.ReadRepoConfig(p.Repository.Name)
    if err != nil {
        return false
    }
    if !p.hasValidEvents(repoConfig) {
        return false
    }
    return true
}

func (p Ping)hasValidEvents(repoConfig domain.RepoConfig) bool {
    events := p.Hook.Events
    found := false
    for _, event := range events {
        for configEvent := range repoConfig.Actions {
            if configEvent == event {
                found = true
                break
            }
        }
        if !found {
            return false    //  --> the ping announces an event that is not
                            //  recognized by the config
        }
    }
    return true
}
