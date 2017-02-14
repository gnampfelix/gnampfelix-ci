package events

import (
    "github.com/gnampfelix/gnampfelix-ci/domain"
    "encoding/json"
    "errors"
)

type PullRequest struct {
    CompareBranch string
    Event
}

//  Push implements Unmarshaler. This is due to the fact that golang currently
//  does not support "required" fields for Unmarshal. A push needs a repository
//  and a ref!
func (p *PullRequest)UnmarshalJSON(data []byte) error {
    var resultMap map[string]interface{}
    err := json.Unmarshal(data, &resultMap)
    if err != nil {
        return err
    }

    action, ok := resultMap["action"].(string)
    if !ok || action != "opened" {
        return errors.New("The pull request could not be read, missing or unsupported action.")
    }

    pullRequest, ok := resultMap["pull_request"].(map[string]interface{})
    if !ok {
        return errors.New("The pull request could not be read, missing pull request data.")
    }
    compareBranch, ok := pullRequest["head"].(map[string]interface{})
    if !ok {
        return errors.New("The pull request could not be read, missing head data.")
    }

    p.CompareBranch, ok = compareBranch["ref"].(string)
    if !ok {
        return errors.New("The pull request could not be read, missing head ref.")
    }
    sha, ok := compareBranch["sha"].(string)
    if !ok {
        return errors.New("The pull request could not be read, missing head sha.")
    }

    compareRepo, ok := compareBranch["repo"].(map[string]interface{})
    if !ok {
        return errors.New("The pull request could not be read, missing head repo data.")
    }
    repoId, ok := compareRepo["id"].(float64)
    if !ok {
        return errors.New("The pull request could not be read, missing head repo id.")
    }

    baseBranch, ok := pullRequest["base"].(map[string]interface{})
    if !ok {
        return errors.New("The pull request could not be read, missing base branch data.")
    }

    ref, ok := baseBranch["ref"].(string) //TODO: Extract only the branch name!
    if !ok {
        return errors.New("The push could not be read, missing base ref.")
    }

    repoMap, ok := baseBranch["repo"].(map[string]interface{})
    if !ok {
        return errors.New("The push could not be read, missing repository information.")
    }

    baseId, ok := repoMap["id"].(float64)
    if !ok {
        return errors.New("The push could not be read, missing base repository id.")
    }

    if repoId != baseId {
        return errors.New("Currently, pull requests from other repositories are not supported.")
    }

    repository, err := domain.CreateRepositoryFromMap(repoMap)
    if err != nil {
        return err
    }

    event, err := newEvent(ref, sha, "push", repository)
    if err != nil {
        return err
    }
    p.Event = event
    return nil
}

func (p PullRequest)HandleEvent() error {
    p.Event.connectWithGitHub()
    defer p.Event.cleanUp()
    p.Event.postPending()

    err := p.Event.cloneRepoAndCheckout()
    if err != nil {
        p.Event.postFailure()
        return err
    }

    err = p.Event.merge(p.CompareBranch)
    if err != nil {
        p.Event.postFailure()
        return err
    }

    err = p.Event.runAction()
    if err != nil {
        p.Event.postFailure()
        return err
    }

    p.Event.postSuccess()
    return err
}
