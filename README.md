# gnampfelix-ci
[![Build Status](http://ci.gnampfelix.de/api/badges/gnampfelix/gnampfelix-ci/status.svg)](http://ci.gnampfelix.de/gnampfelix/gnampfelix-ci)

gnampfelix-ci is a continous integration platform that can be hosted on any linux system. It handles incoming requests from the GitHub Webhook API and executes test scripts as you define them

## 1. Installation
Download this repository with `go get github.com/gnampfelix/gnampfelix-ci`. This will download gnampfelix-ci as well as all dependencies. Run `go install github.com/gnampfelix/gnampfelix-ci`.

## 2. Configuration
gnampfelix-ci awaits a configuration file called "gnampfile". This file can contain an empty json object ({}), but I do not recommend that.
```     
{
    "PreventHTTPS": false,          
    "Port": 8080,
    "Certificate": "certificate.pem",
    "Keyfile": "key.ppk",
    "CiRoot": "./",
    "GithubSecret": "abcdefg"
}
```
This demonstrates a complete configuration file. If you want to run gnampfelix ci in HTTPS mode, set "PreventHTTPS" to false. You will need a key file and a certificate file to run in HTTPS mode, so specify their relative paths in "Certificate" and "Keyfile".
If you want to secure your gnampfelix ci installation, set up a "GithubSecret" in order to validate every request. For more information on how this works, have a look at [the GitHub API reference](https://developer.github.com/webhooks/securing/).
Make sure that the port you specified is not already in use and is not blocked by your firewall or router.

## 3. Set up a repository
To make gnampfelix ci handle on notifications on a specific repository, you will need to create a file called "yourRepo.json" in the ci root folder, where yourRepo is the name of your repository.
```
{
    "actions": {
        "push":[
        {
            "Branches": ["master"],
            "PreTest": "pre.sh",
            "Test": "test.sh"
        }
        ],
        "pull_request":[
        {
            "Branches": ["master"],
            "PreTest": "pre.sh",
            "Test": "test.sh"
        }
        ]
    },
    "User": "UserName",
    "Token": "Token"
}
```
This is an example file. For each event you can add as much handlers you want. Each handler must have a "Branches" array, a "PreTest" and a "Test" script. To handle all branches with one handler, use "{all}". If you use "{all}" and specific branch names, the handler connected to the specific branch name will be used. If you have multiple handlers that have the same branch names, the first handler will be used.
You need to set up the username and the access token as well. This is not used for any "git" actions right now, but is needed to update the test status of commits on GitHub. If your repository is private. you need to set up your git with `git config --global credential.helper store`. Before using gnampfelix ci, you must have accessed your private repository at least once with this git setting. (I plan to change this so that gnampfelix ci uses the username and the token in the repo configuration file.)


To get notified whenever something happens on your github repository, set up a webhook for this repoitory. The address to gnampfelix ci depends on where you installed gnampfelix ci and how you configured it. As an example, a gnampfelix ci installation running on a machine that is available under http://example.org that is configured to run on port 8080 will be available under http://example.org:8080/notifications. Make sure to add the path to notifications! If you configured your gnampfelix ci to use a github secret, set your secret here as well. You can chose any webhook configuration you like, but be reminded that gnampfelix ci currently only supports push and pull_request events. After setting up the webhook, github sends a "ping" event. If you configured everything correctly, this ping will not fail.
