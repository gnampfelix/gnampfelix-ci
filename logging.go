package main
import (
    "log"
    "io"
)

var (
    Warning *log.Logger
    Info *log.Logger
    Error *log.Logger
)

func initLogger(warningIO io.Writer, infoIO io.Writer, errorIO io.Writer) {
    Warning = log.New(warningIO, "[WARNING]\t", 0)
    Info = log.New(infoIO, "[INFO]\t\t", 0)
    Error = log.New(errorIO, "[ERROR]\t\t", 0)
}
