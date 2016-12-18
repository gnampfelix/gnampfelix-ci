package domain

import (
    "os/exec"
)

//  Git represents the Git Repository to use for building, testing and deploying.
//  While Repository only holds the abstract Repository data, Git represents the
//  actual git client to execute the typical git actions.
type Git struct {
    User string
    Token string
    Repo Repository
}

func (git *Git)CreateNewGit(User string, Token string, Repo Repository) ([]byte, error) {
    cmd := exec.Command("git", "config --global credential.helper store")
    output, err := cmd.CombinedOutput()
    git.User = User
    git.Token = Token
    git.Repo = Repo
    return output, err
}

//  Clone the specified git repository. The output and any errors of this action
//  are returned.
func (git Git)Clone() ([]byte, error) {
    cmd := exec.Command("git", "clone", git.Repo.CloneURL)
    output, err := cmd.CombinedOutput()
    return output, err
}

//  Remove the git repository from the file system. The output and any errors of
//  actions are returned.
func (git Git)Remove() ([]byte, error) {
    cmd := exec.Command("rm", "-rf", git.Repo.Name)
    output, err := cmd.CombinedOutput()
    return output, err
}
