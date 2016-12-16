package middleware

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
)

func HandleIncomingNotification(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    //Convert request body into map
    var payload map[string]interface{}
    decoder := json.NewDecoder(request.Body)
    defer request.Body.Close()

    err := decoder.Decode(&payload)
    if err != nil {
        warningLog.Println(err)
        http.NotFound(writer, request)
    }

    notificationType := request.Header.Get("X-GitHub-Event")
    switch notificationType {
    case "ping":
        handlePing(writer, payload)
    case "push":
        handlePush(writer, payload)
    default:
        http.NotFound(writer, request)
    }

}

func handlePing(writer http.ResponseWriter, payload map[string]interface{}) {
    repository, ok := payload["repository"].(map[string]interface{})
    if !ok {
        infoLog.Printf("Invalid ping request recieved.\n")
        http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }

    repoName, ok := repository["name"].(string)
    if !ok {
        infoLog.Printf("Invalid ping request recieved.\n")
        http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }
    infoLog.Printf("Incoming webhook ping from %s, starting verification..\n", repoName)
    writer.WriteHeader(http.StatusAccepted)
}

func handlePush(writer http.ResponseWriter, payload map[string]interface{}) {

}
