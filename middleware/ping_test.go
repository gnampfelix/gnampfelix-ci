package middleware

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
    "github.com/julienschmidt/httprouter"
)

func TestValidPing(t *testing.T) {
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

func TestInvalidPing(t *testing.T) {
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
            status, http.StatusOK)
    }
}

func buildRouter() *httprouter.Router {
    notificationRouter := NewRouter()
    notificationRouter.POST("/notifications", HandleIncomingNotification)
    return notificationRouter
}
