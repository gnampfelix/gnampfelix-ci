package middleware

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
)

func TestValidPush(t *testing.T) {
    prepareTestEnvironment()
    postPayload := strings.NewReader(notificationPushValid)
    req, err := http.NewRequest("POST", "/notifications", postPayload)
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("X-GitHub-Event", "push")

    rr := httptest.NewRecorder()
    notificationRouter := buildRouter()
    notificationRouter.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusAccepted {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}
