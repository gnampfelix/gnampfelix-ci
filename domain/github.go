package domain

import (
    "net/http"
    "io"
    "encoding/json"
    "time"
    "bytes"
)

//  GitHub represents the GitHub Remote. Therefore, GitHub wraps Git and adds further
//  functionality which is designed to communicate with github.com.
type GitHub struct {
    Git
}

type GitHubStatusError int
func (s GitHubStatusError)Error() string {
    return "The status could not be posted to GitHub."
}

var client = http.Client{Timeout: time.Second * 10}

//  Post the given Status and the given sha to the GitHub API.
func (g *GitHub)PostStatus(status Status, sha string) error {
    url := "https://api.github.com/repos/" + g.Repo.Owner + "/" + g.Repo.Name + "/statuses/" + sha
    payload, err := json.Marshal(status)
    if err != nil {
        return err
    }
    err = g.post(bytes.NewReader(payload), url)
    return err
}


func (g *GitHub)post(body io.Reader, url string) error {
    request, err := http.NewRequest(http.MethodPost, url, body)
    if err != nil {
        return err
    }
    request.Header.Set("Authorization", "token " + g.Token)

    response, err := client.Do(request)
    if err != nil {
        return err
    }
    defer response.Body.Close()
    if response.StatusCode != http.StatusCreated {
        return GitHubStatusError(1)
    }
    return nil
}
