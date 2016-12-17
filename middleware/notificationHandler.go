package middleware

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/Thamtham/gnampfelix-ci/events"
    "encoding/json"
)

func HandleIncomingNotification(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    decoder := json.NewDecoder(request.Body)
    defer request.Body.Close()


    notificationType := request.Header.Get("X-GitHub-Event")
    switch notificationType {
    case "ping":
        var ping events.Ping
        err := decoder.Decode(&ping)
        if err != nil {
            errorLog.Println(err)
            http.Error(writer, err.Error(), http.StatusBadRequest)
            return
        }
        infoLog.Printf("Incoming ping for %s, starting verification.\n", ping.Repository.Name)
        if !ping.HasValidConfiguration() {
            errorLog.Printf("The ping for %s does not have a valid configuration file.\n", ping.Repository.Name)
            http.Error(writer, "No valid configuration for this repository found.", http.StatusNotAcceptable)
            return
        }
        writer.WriteHeader(http.StatusAccepted)
        return
    case "push":
        var push events.Push
        err := decoder.Decode(&push)
        if err != nil {
            errorLog.Println(err)
            http.Error(writer, err.Error(), http.StatusBadRequest)
            return
        }
        infoLog.Printf("Incoming push event for %s, starting build run.\n", push.Repository.Name)
        err = push.HandleEvent()
        if err != nil {
            errorLog.Println(err)
        }
        writer.WriteHeader(http.StatusAccepted)
        return
    default:
        http.NotFound(writer, request)
    }

}

func handlePush(writer http.ResponseWriter, payload map[string]interface{}) {

}
