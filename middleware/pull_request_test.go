package middleware

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
)

func TestValidPullRequest(t *testing.T) {
    prepareTestEnvironment()
    postPayload := strings.NewReader(notificationPullRequestValid)
    req, err := http.NewRequest("POST", "/notifications", postPayload)
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("X-GitHub-Event", "pull_request")
    rr := httptest.NewRecorder()
    notificationRouter := buildRouter()
    notificationRouter.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusAccepted {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}
