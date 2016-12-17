package domain

type RepoConfig struct {
    Actions map[string][]Action
}

type Action struct {
    Branches []string
    PreTest string
    Test string
    Deploy Deployment
}
