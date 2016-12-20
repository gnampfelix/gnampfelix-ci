package domain

//  The GitHub status of a commit.
type Status struct {
    State string `json:"state"`
    TargetURL string `json:"target_url"`
    Description string `json:"description"`
    Context string `json:"context"`
}

//  Creates the status "pending".
func (s *Status)Pending(targetUrl string) Status{
    s.State = "pending"
    s.Description = "The build is currently in progress."
    s.setCommonAttributes(targetUrl)
    return *s
}

//  Creates the status "success".
func (s *Status)Success(targetUrl string) Status{
    s.State = "success"
    s.Description = "The build succeeded!"
    s.setCommonAttributes(targetUrl)
    return *s
}

//  Creates the status "error".
func (s *Status)Error(targetUrl string) Status{
    s.State = "error"
    s.Description = "There was an error during the build. The build was not executed."
    s.setCommonAttributes(targetUrl)
    return *s
}

//  Creates the status "failure".
func (s *Status)Failure(targetUrl string) Status{
    s.State = "failure"
    s.Description = "The build/test failed. Please visit gci for further inforation."
    s.setCommonAttributes(targetUrl)
    return *s
}

func (s *Status)setCommonAttributes(targetUrl string) {
    s.TargetURL = targetUrl
    s.Context = "continuous-integration/gci"
}
