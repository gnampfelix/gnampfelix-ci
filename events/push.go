package events

import (
    "github.com/Thamtham/gnampfelix-ci/domain"
    "encoding/json"
    "errors"
)

//  GitHub sends a push notification after each push on the repo the webhook is
//  assigned to. The type Push represents this notification for further computation.
type Push struct {
    Event
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

    ref, ok := resultMap["ref"].(string) //TODO: Extract only the branch name!
    if !ok {
        return errors.New("The push could not be read, missing ref.")
    }

    sha, ok := resultMap["after"].(string)
    if !ok {
        return errors.New("The push could not be read, missing sha (\"after\").")
    }

    repoMap, ok := resultMap["repository"].(map[string]interface{})
    if !ok {
        return errors.New("The push could not be read, missing repository information.")
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

//  Handle the Push event:
//  Find the matching action for this event in the configFile, clone the repo
//  execute the actions and remove the repo again. The output of this steps
//  will be saved in a file. The file is closed afterwards.
func (p Push)HandleEvent() error {
    p.Event.connectWithGitHub()
    defer p.Event.cleanUp()
    p.Event.postPending()

    err := p.Event.cloneRepoAndCheckout()
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
