package domain

import (
    "errors"
)

//  Hook represents all the necessary data that is send as the "hook" part
//  of the GitHub ping.
type Hook struct {
    Events []string `json:"events"`
}

func CreateHookFromMap(data map[string]interface{}) (Hook, error) {
    eventsInterface, ok := data["events"].([]interface{})
    if !ok {
        return Hook{}, errors.New("Can't read hook data, missing events.")
    }

    events := make([]string, len(eventsInterface))
    for i := range events {
        events[i] = eventsInterface[i].(string)
    }

    return Hook{Events: events}, nil
}
