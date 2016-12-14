package main

import (
    "net/http"
    "os"
    "strconv"
)

func main() {
    notificationRouter := NewRouter()
    initLogger(os.Stdout, os.Stdout, os.Stdout)
    //notificationRouter.POST("/notifications", handleIncomingNotification)

    middleware := Middleware{}
    middleware.Add(notificationRouter)

    config, err := readConfing()
    if err != nil {
        Error.Fatal(err)
    }

    if config.PreventHTTPS {
        err = http.ListenAndServe(":" + strconv.Itoa(config.Port), nil)
    } else {
        err = http.ListenAndServeTLS(":" + strconv.Itoa(config.Port), config.Certificate, config.Keyfile, nil)
    }
    if err != nil {
        Error.Fatal(err)
    }

}
