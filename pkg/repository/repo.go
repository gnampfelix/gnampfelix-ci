package repository

import (
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	gitHttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"os"
)

type Repository interface {
	SetId(buildId string)
	SetUser(user string)
	SetToken(token string)
	SetUrl(url string)
	SetName(name string)
	SetRef(ref string)
	EnsureLocalExistence() error
	ParseConfig() (RepoConfig, error)
	Delete()
}

func New() Repository {
	return &repository{}
}

type repository struct {
	repoCloner
	repoDeleter
	repoParser
}

func (r *repository) SetId(buildId string) {
	r.repoCloner.id = buildId
	r.repoDeleter.id = buildId
	r.repoParser.id = buildId
}

func (r *repository) SetUser(user string) {
	r.repoCloner.user = user
}

func (r *repository) SetToken(token string) {
	r.repoCloner.token = token
}

func (r *repository) SetUrl(url string) {
	r.repoCloner.url = url
}

func (r *repository) SetName(name string) {
	r.repoCloner.name = name
	r.repoParser.name = name
}

func (r *repository) SetRef(ref string) {
	r.repoCloner.ref = ref
}

type repoCloner struct {
	id    string
	user  string
	token string
	url   string
	name  string
	ref   string
}

func (r *repoCloner) EnsureLocalExistence() error {
	auth := gitHttp.NewBasicAuth(r.user, r.token)
	repository, err := git.PlainClone(r.id+"/"+r.name, false, &git.CloneOptions{
		URL:  r.url,
		Auth: auth,
	})
	if err != nil {
		return err
	}

	wt, err := repository.Worktree()
	if err != nil {
		return err
	}

	err = wt.Checkout(&git.CheckoutOptions{Branch: plumbing.ReferenceName(r.ref)})
	return err
}

type repoDeleter struct {
	id string
}

func (r *repoDeleter) Delete() {
	os.RemoveAll(r.id)
}

type repoParser struct {
	id   string
	name string
}

func (r *repoParser) ParseConfig() (RepoConfig, error) {
	return Parse(r.id + "/" + r.name + "/gnampfile.json")
}
