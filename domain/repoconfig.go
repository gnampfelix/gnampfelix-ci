package domain

import (
    "errors"
    "os/exec"
)

//  A RepoConfig is the configuration file for a repository in gci.
//      {
//          "actions":{
//              "push":[
//                  {
//                      "Branches": [...],
//                      "PreTest": "xyz.sh",
//                      "Test": "abc.sh",
//                      "Deployment": currentlyNotSupported
//                  },
//                  {
//                      "Branches": [...],
//                      "PreTest": "xyz.sh",
//                      "Test": "abc.sh",
//                      "Deployment": currentlyNotSupported
//                  },
//              ],
//              "pull_request":[],
//          },
//          "user":"UserName",
//          "token": "token"
//      }
type RepoConfig struct {
    Actions map[string][]Action `json:"actions"`
    Username string `json:"user"`
    AccessToken string `json:"token"`
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
