package main

import (
    "net/http"
    "os"
    "strconv"
    "github.com/gnampfelix/gnampfelix-ci/middleware"
    "github.com/gnampfelix/gnampfelix-ci/config"
)

func main() {
    initLogger(os.Stdout, os.Stdout, os.Stdout)
    notificationRouter := middleware.NewRouter()
    notificationRouter.POST("/notifications", middleware.HandleIncomingNotification)

    middleware := middleware.New()
    middleware.Add(notificationRouter)
    config.SetLogger(Error, Warning, Info)
    config, err := config.ReadFile()
    if err != nil {
        Error.Fatal(err)
    }

    if config.PreventHTTPS {
        err = http.ListenAndServe(":" + strconv.Itoa(config.Port), middleware)
    } else {
        err = http.ListenAndServeTLS(":" + strconv.Itoa(config.Port), config.Certificate, config.Keyfile, middleware)
    }
    if err != nil {
        Error.Fatal(err)
    }

}
