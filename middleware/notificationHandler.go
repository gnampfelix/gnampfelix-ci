package middleware

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/gnampfelix/gnampfelix-ci/events"
    "encoding/json"
)

func HandleIncomingNotification(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    decoder := json.NewDecoder(request.Body)
    defer request.Body.Close()
    notificationType := request.Header.Get("X-GitHub-Event")
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
    infoLog.Printf("Incoming pull_request event for %s, starting build run.\n", pullRequest.Repository.Name)
    err = pullRequest.HandleEvent()
    if err != nil {
        errorLog.Println(err)
    }
    writer.WriteHeader(http.StatusAccepted)
}

func printAndWriteError(writer http.ResponseWriter, err error) {
    errorLog.Println(err)
    http.Error(writer, err.Error(), http.StatusBadRequest)
}
