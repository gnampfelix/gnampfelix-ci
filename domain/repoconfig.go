package domain

import (
    "errors"
    "os/exec"
)

//  A RepoConfig is the configuration file for a repository in gci.
//      {
//          "push":[
//              {
//                  "Branches": [...],
//                  "PreTest": "xyz.sh",
//                  "Test": "abc.sh",
//                  "Deployment": currentlyNotSupported
//              },
//              {
//                  "Branches": [...],
//                  "PreTest": "xyz.sh",
//                  "Test": "abc.sh",
//                  "Deployment": currentlyNotSupported
//              },
//          ],
//          "pull_request":[]
//      }
type RepoConfig struct {
    Actions map[string][]Action
}

//  With the given git ref, GetAction returns the action first action that has
//  ref in its Branches. If no Action is found and if present, the default action
//  is selected. The event specifies for which GitHub event the action is, eg "push".
func (r *RepoConfig)GetAction(ref string, event string) (Action, error) {
    actions, ok := r.Actions[event]
    if !ok {
        return Action{}, errors.New("The config does not specify an event for " + event + ".")
    }

    var defaultAction Action
    foundDefault := false
    for _, action := range actions {
        for _, branch := range action.Branches {
            if branch == "{all}" {
                defaultAction = action
                foundDefault = true
            } else if branch == ref {
                return action, nil
            }
        }
    }
    if foundDefault {
        return defaultAction, nil
    }
    return Action{}, errors.New("The config does not specify an action for this event and this ref. No default found.")
}

type Action struct {
    Branches []string
    PreTest string
    Test string
    Deploy Deployment
}

func (a Action)Run() ([]byte, error) {
    cmd := exec.Command(a.PreTest)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return output, err
    }

    cmd = exec.Command(a.Test)
    output, err = cmd.CombinedOutput()
    if err != nil {
        return output, err
    }
    return output, nil
}
