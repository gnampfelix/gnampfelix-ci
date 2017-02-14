package events

import (
    "github.com/gnampfelix/gnampfelix-ci/domain"
    "github.com/gnampfelix/gnampfelix-ci/config"
    "os"
    "time"
)

var status domain.Status

type Event struct {
    Ref string
    Sha string
    Repository domain.Repository
    gitHub domain.GitHub
    repoConfig domain.RepoConfig
    mainConfig config.ConfigFile
    action domain.Action
    logFile *os.File
    isLogStd bool
}

func newEvent(ref, sha, actionType string, repo domain.Repository) (Event, error) {
    mainConfig := config.GetConfig()
    repoConfig, err := config.ReadRepoConfig(repo.Name)
    if err != nil {
        return Event{}, err
    }
    fileName := time.Now().Format("02-01-2006_15-04-05") + ".log"
    isLogStd := false
    logFile, err := os.Create(fileName)
    if err != nil {
        logFile = os.Stdout
        isLogStd = true
    }
    action, err := repoConfig.GetAction(ref, actionType)
    return Event{Ref:ref, Sha:sha,
        Repository:repo,
        repoConfig: repoConfig,
        mainConfig: mainConfig,
        action: action,
        logFile: logFile,
        isLogStd: isLogStd,
        }, err
}

func (e *Event)connectWithGitHub() {
    var git domain.Git
    git.CreateNewGit(e.repoConfig.Username, e.repoConfig.AccessToken, e.Repository)
    e.gitHub = domain.GitHub{Git:git}
}

func (e Event)cloneRepoAndCheckout() error {
    cloneResult, err := e.gitHub.Git.Clone()
    e.logFile.Write(cloneResult)
    if err != nil {
        return err
    }
    checkoutResult, err := e.gitHub.Git.Checkout(e.Ref)
    e.logFile.Write(checkoutResult)
    return err
}

func (e Event)merge(compareBranch string) error {
    //need to checkout first to make the branch known to git
    checkoutResult, err := e.gitHub.Git.Checkout(compareBranch)
    e.logFile.Write(checkoutResult)
    if err != nil {
        return err
    }
    //Checkout back..
    checkoutResult, err = e.gitHub.Git.Checkout(e.Ref)
    e.logFile.Write(checkoutResult)
    if err != nil {
        return err
    }
    mergeResult, err := e.gitHub.Git.Merge(compareBranch)
    e.logFile.Write(mergeResult)
    return err
}

func (e Event)runAction() error {
    actionResult, err := e.action.Run(e.mainConfig.CiRoot)
    e.logFile.Write(actionResult)
    return err
}

func (e Event)postPending() {
    err := e.gitHub.PostStatus(status.Pending("http://google.de"), e.Sha)
    if err != nil {
        e.logFile.WriteString(err.Error())
    }
}

func (e Event)postFailure() {
    err := e.gitHub.PostStatus(status.Failure("http://google.de"), e.Sha)
    if err != nil {
        e.logFile.WriteString(err.Error())
    }
}

func (e Event)postSuccess() {
    err := e.gitHub.PostStatus(status.Success("http://google.de"), e.Sha)
    if err != nil {
        e.logFile.WriteString(err.Error())
    }
}

func (e Event)cleanUp() {
    e.gitHub.Git.Remove()
    if !e.isLogStd {
        e.logFile.Close()
    }
}
