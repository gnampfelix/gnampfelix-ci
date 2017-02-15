package middleware

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/gnampfelix/gnampfelix-ci/events"
    "encoding/json"
    "io"
    "crypto/hmac"
    "crypto/sha1"
    "hash"
    "encoding/hex"
)

var hasher hash.Hash
var secret string
var GithubSecret = ""

func SetGithubSecret(newSecret string) {
    GithubSecret = newSecret
}

func HandleIncomingNotification(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    hasher = hmac.New(sha1.New, []byte(GithubSecret))
    reader := io.TeeReader(request.Body, hasher)
    decoder := json.NewDecoder(reader)
    defer request.Body.Close()
    notificationType := request.Header.Get("X-GitHub-Event")
    secret = request.Header.Get("X-Hub-Signature")
    switch notificationType {
    case "ping":
        handlePing(writer, decoder)
        return
    case "push":
        handlePush(writer, decoder)
        return
    case "pull_request":
        handlePullRequest(writer, decoder)
        return;
    default:
        http.NotFound(writer, request)
    }
}

func handlePing(writer http.ResponseWriter, decoder *json.Decoder) {
    var ping events.Ping
    err := decoder.Decode(&ping)
    if err != nil {
        printAndWriteError(writer, err)
        return
    }
    if !isAuthorized() {
        http.Error(writer, "Invalid signature", http.StatusForbidden)
        warningLog.Printf("Incoming ping event for %s with invalid signature.\n", ping.Repository.Name)
        return;
    }
    infoLog.Printf("Incoming ping for %s, starting verification.\n", ping.Repository.Name)
    if !ping.HasValidConfiguration() {
        errorLog.Printf("The ping for %s does not have a valid configuration file.\n", ping.Repository.Name)
        http.Error(writer, "No valid configuration for this repository found.", http.StatusNotAcceptable)
        return
    }
    writer.WriteHeader(http.StatusAccepted)
}

func handlePush(writer http.ResponseWriter, decoder *json.Decoder) {
    var push events.Push
    err := decoder.Decode(&push)
    if err != nil {
        printAndWriteError(writer, err)
        return
    }
    if !isAuthorized() {
        http.Error(writer, "Invalid signature", http.StatusForbidden)
        warningLog.Printf("Incoming push event for %s with invalid signature.\n", push.Repository.Name)
        return;
    }
    infoLog.Printf("Incoming push event for %s, starting build run.\n", push.Repository.Name)
    err = push.HandleEvent()
    if err != nil {
        errorLog.Println(err)
    }
    writer.WriteHeader(http.StatusAccepted)
}

func handlePullRequest(writer http.ResponseWriter, decoder *json.Decoder) {
    var pullRequest events.PullRequest
    err := decoder.Decode(&pullRequest)
    if err != nil {
        printAndWriteError(writer, err)
        return
    }
    if !isAuthorized() {
        http.Error(writer, "Invalid signature", http.StatusForbidden)
        warningLog.Printf("Incoming pull_request event for %s with invalid signature.\n", pullRequest.Repository.Name)
        return;
    }
    infoLog.Printf("Incoming pull_request event for %s, starting build run.\n", pullRequest.Repository.Name)
    err = pullRequest.HandleEvent()
    if err != nil {
        errorLog.Println(err)
    }
    writer.WriteHeader(http.StatusAccepted)
}

func isAuthorized() bool {
    if GithubSecret == "" {
        return true
    }
    expectedSignature := hasher.Sum(nil)
    signature := "sha1=" + hex.EncodeToString(expectedSignature[:])
    return hmac.Equal([]byte(signature), []byte(secret))
}

func printAndWriteError(writer http.ResponseWriter, err error) {
    errorLog.Println(err)
    if isAuthorized() {
        http.Error(writer, err.Error(), http.StatusBadRequest)
    } else {
        http.Error(writer, "Invalid signature", http.StatusForbidden)
    }
}
