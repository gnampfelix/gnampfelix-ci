package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
)

func handleIncomingNotification(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    //Convert request body into map
    var payload map[string]interface{}
    decoder := json.NewDecoder(request.Body)
    defer request.Body.Close()

    err := decoder.Decode(payload)
    if err != nil {
        Warning.Println(err)
        http.NotFound(writer, request)
    }
    Info.Printf("Incoming event on %s.\n", payload["ref"].(string))
    writer.WriteHeader(http.StatusAccepted)
}
