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
            warningLog.Println(err)
            http.Error(writer, err.Error(), http.StatusBadRequest)
            return
        }
        writer.WriteHeader(http.StatusAccepted)
        return
    case "push":

    default:
        http.NotFound(writer, request)
    }

}

func handlePush(writer http.ResponseWriter, payload map[string]interface{}) {

}
