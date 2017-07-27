package repository

import (
	"encoding/json"
	"errors"
	"github.com/gnampfelix/gnampfelix-ci/pkg/infrastructure"
	"os"
)

type RepoConfig interface {
	GetAction(ref string) (Action, error)
}

func Parse(path string) (RepoConfig, error) {
	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(configFile)

	buffer := make(map[string]interface{})
	err = decoder.Decode(&buffer)
	if err != nil {
		return nil, err
	}

	result := repoConfig{
		tasks:   make(map[string]Task),
		actions: make(map[string]Action),
	}

	tasks, ok := buffer["tasks"].(map[string]interface{})
	if !ok {
		return nil, errors.New("No \"tasks\" section in the gnampfile found.")
	}
	actions, ok := buffer["actions"].(map[string]interface{})
	if !ok {
		return nil, errors.New("No \"actions\" section in the gnampfile found.")
	}

	for k := range tasks {
		task, err := createTask(tasks[k].(map[string]interface{}), k)
		if err != nil {
			return nil, err
		}
		result.tasks[k] = task
	}

	for k := range actions {
		action, err := createAction(actions[k].(map[string]interface{}), k, result.tasks, path)
		if err != nil {
			return nil, err
		}
		result.actions[k] = action
	}
	return result, nil
}

func createTask(input map[string]interface{}, id string) (task, error) {
	pre, ok := input["pre"].(string)
	if !ok {
		return task{}, errors.New("The task with the id \"" + id + "\" does not provide a valid \"pre\"-field.")
	}
	test, ok := input["test"].(string)
	if !ok {
		return task{}, errors.New("The task with the id \"" + id + "\" does not provide a valid \"test\"-field.")
	}
	post, ok := input["post"].(string)
	if !ok {
		return task{}, errors.New("The task with the id \"" + id + "\" does not provide a valid \"post\"-field.")
	}
	return task{
		pre:  pre,
		test: test,
		post: post,
		id:   id,
	}, nil
}

func createAction(input map[string]interface{}, ref string, tasks map[string]Task, path string) (action, error) {
	envId, ok := input["env"].(string)
	if !ok {
		return action{}, errors.New("The action for the ref \"" + ref + "\" does not a valid \"env\"-field.")
	}

	taskId, ok := input["task"].(string)
	if !ok {
		return action{}, errors.New("The action for the ref \"" + ref + "\" does not a valid \"task\"-field.")
	}

	task, ok := tasks[taskId]
	if !ok {
		return action{}, errors.New("The task \"" + taskId + "\" used in action for the ref \"" + ref + "\" does not exist.")
	}

	env, err := infrastructure.GetEnvironment(envId, path)
	if err != nil {
		return action{}, err
	}

	return action{
		ref:  ref,
		task: task,
		env:  env,
	}, nil
}

type repoConfig struct {
	tasks   map[string]Task
	actions map[string]Action
}

func (r repoConfig) GetAction(ref string) (Action, error) {
	action, ok := r.actions[ref]
	if ok {
		return action, nil
	}
	action, ok = r.actions["{{ all }}"]
	if ok {
		return action, nil
	}
	return nil, errors.New("Action for ref \"" + ref + "\" not found.")
}

type Task interface {
	GetPre() string
	GetTest() string
	GetPost() string
}

type task struct {
	id   string
	pre  string
	test string
	post string
}

func (t task) GetPre() string {
	return t.pre
}

func (t task) GetTest() string {
	return t.test
}

func (t task) GetPost() string {
	return t.post
}

type Action interface {
	GetTask() Task
	GetEnvironment() infrastructure.BuildEnvironment
}

type action struct {
	ref  string
	task Task
	env  infrastructure.BuildEnvironment
}

func (a action) GetTask() Task {
	return a.task
}

func (a action) GetEnvironment() infrastructure.BuildEnvironment {
	return a.env
}
