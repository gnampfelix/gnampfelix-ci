package domain

import (
    "errors"
)

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
