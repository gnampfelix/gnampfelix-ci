package events

import (
    "github.com/Thamtham/gnampfelix-ci/domain"
    "github.com/Thamtham/gnampfelix-ci/config"
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "time"
)

//  GitHub sends a push notification after each push on the repo the webhook is
//  assigned to. The type Push represents this notification for further computation.
type Push struct {
    Ref string
    Sha string
    Repository domain.Repository
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

    ref, ok := resultMap["ref"].(string)
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

    p.Ref = ref
    p.Repository = repository
    p.Sha = sha
    return nil
}

//  Handle the Push event:
//  Find the matching action for this event in the configFile, clone the repo
//  execute the actions and remove the repo again. The output of this steps
//  will be saved in a file.
func (p Push)HandleEvent() error {
    mainConfig := config.GetConfig()
    var (
        git domain.Git
        gitHub domain.GitHub
        status domain.Status
    )

    repoConfig, err := config.ReadRepoConfig(p.Repository.Name)
    if err != nil {
        return err
    }
    action, err := repoConfig.GetAction(p.Ref, "push")
    if err != nil {
        return err
    }

    git.CreateNewGit(repoConfig.Username, repoConfig.AccessToken, p.Repository)
    gitHub = domain.GitHub{Git:git}

    p.postStatusAndPrint(gitHub, status.Pending("http://google.de"))

    fileName := time.Now().Format("02-01-2006_15-04-05") + ".log"
    logFile, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer logFile.Close()

    gitResult, err := git.Clone()
    logFile.Write(gitResult)
    if err != nil {
        p.postStatusAndPrint(gitHub, status.Error("http://google.de"))
        return err
    }

    actionOutput, err := action.Run(mainConfig.CiRoot)
    logFile.Write(actionOutput)
    if err != nil {
        p.postStatusAndPrint(gitHub, status.Error("http://google.de"))
        return err
    }

    gitResult, err = git.Remove()
    logFile.Write(gitResult)
    if err != nil {
        p.postStatusAndPrint(gitHub, status.Error("http://google.de"))
        return err
    }
    p.postStatusAndPrint(gitHub, status.Success("http://google.de"))
    return err
}

func (p Push)postStatusAndPrint(gitHub domain.GitHub, status domain.Status) {
    err := gitHub.PostStatus(status, p.Sha)
    if err != nil {
        fmt.Println(err)
    }
}
