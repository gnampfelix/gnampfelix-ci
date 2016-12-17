package middleware

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
    "github.com/julienschmidt/httprouter"
    "io/ioutil"
)

func TestValidPing(t *testing.T) {
    saveConfigurationFile()
    postPayload := strings.NewReader(notificationPingValid)
    req, err := http.NewRequest("POST", "/notifications", postPayload)
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("X-GitHub-Event", "ping")

    rr := httptest.NewRecorder()
    notificationRouter := buildRouter()
    notificationRouter.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusAccepted {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}

func TestValidPingInvalidConfig(t *testing.T) {
    saveConfigurationFile()
    postPayload := strings.NewReader(notificationPingValidInvalidConfig)
    req, err := http.NewRequest("POST", "/notifications", postPayload)
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("X-GitHub-Event", "ping")

    rr := httptest.NewRecorder()
    notificationRouter := buildRouter()
    notificationRouter.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusNotAcceptable {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusNotAcceptable)
    }
}

func TestInvalidPing(t *testing.T) {
    saveConfigurationFile()
    postPayload := strings.NewReader(notificationPingInvalid)
    req, err := http.NewRequest("POST", "/notifications", postPayload)
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("X-GitHub-Event", "ping")

    rr := httptest.NewRecorder()
    notificationRouter := buildRouter()
    notificationRouter.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }
}

func buildRouter() *httprouter.Router {
    notificationRouter := NewRouter()
    notificationRouter.POST("/notifications", HandleIncomingNotification)
    return notificationRouter
}

func saveConfigurationFile() {
    ioutil.WriteFile("gnampfelix-ci-tests.json", []byte(emptyConfigFile), 0)
}
