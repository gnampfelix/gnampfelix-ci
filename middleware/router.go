package middleware

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "os"
)

var (
    infoLog *log.Logger
    warningLog *log.Logger
    errorLog *log.Logger
)

func init() {
    infoLog = log.New(os.Stdout, "[INFO]\t\t", 0)
    warningLog = log.New(os.Stdout, "[WARNING]\t\t", 0)
    errorLog = log.New(os.Stdout, "[ERROR]\t\t", 0)
}

type NotFound struct{}

func (n *NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	notFound := new(NotFound)
	router.NotFound = notFound
	return router
}
