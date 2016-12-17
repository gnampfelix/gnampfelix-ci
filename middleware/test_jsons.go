package middleware

var notificationPingValid = `
{
      "zen": "Encourage flow.",
      "hook_id": 11169236,
      "hook": {
        "type": "Repository",
        "id": 11169236,
        "name": "web",
        "active": true,
        "events": [
          "pull_request",
          "push"
        ],
        "config": {
          "content_type": "json",
          "insecure_ssl": "0",
          "url": "http://gnampfelix.de:8080/notification"
        },
        "updated_at": "2016-12-15T18:00:00Z",
        "created_at": "2016-12-15T18:00:00Z",
        "url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/hooks/11169236",
        "test_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/hooks/11169236/test",
        "ping_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/hooks/11169236/pings",
        "last_response": {
          "code": null,
          "status": "unused",
          "message": null
        }
      },
      "repository": {
        "id": 76581419,
        "name": "gnampfelix-ci-tests",
        "full_name": "Thamtham/gnampfelix-ci-tests",
        "owner": {
          "login": "Thamtham",
          "id": 10077533,
          "avatar_url": "https://avatars.githubusercontent.com/u/10077533?v=3",
          "gravatar_id": "",
          "url": "https://api.github.com/users/Thamtham",
          "html_url": "https://github.com/Thamtham",
          "followers_url": "https://api.github.com/users/Thamtham/followers",
          "following_url": "https://api.github.com/users/Thamtham/following{/other_user}",
          "gists_url": "https://api.github.com/users/Thamtham/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/Thamtham/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/Thamtham/subscriptions",
          "organizations_url": "https://api.github.com/users/Thamtham/orgs",
          "repos_url": "https://api.github.com/users/Thamtham/repos",
          "events_url": "https://api.github.com/users/Thamtham/events{/privacy}",
          "received_events_url": "https://api.github.com/users/Thamtham/received_events",
          "type": "User",
          "site_admin": false
        },
        "private": true,
        "html_url": "https://github.com/Thamtham/gnampfelix-ci-tests",
        "description": null,
        "fork": false,
        "url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests",
        "forks_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/forks",
        "keys_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/keys{/key_id}",
        "collaborators_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/collaborators{/collaborator}",
        "teams_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/teams",
        "hooks_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/hooks",
        "issue_events_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/issues/events{/number}",
        "events_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/events",
        "assignees_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/assignees{/user}",
        "branches_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/branches{/branch}",
        "tags_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/tags",
        "blobs_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/blobs{/sha}",
        "git_tags_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/tags{/sha}",
        "git_refs_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/refs{/sha}",
        "trees_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/trees{/sha}",
        "statuses_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/statuses/{sha}",
        "languages_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/languages",
        "stargazers_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/stargazers",
        "contributors_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/contributors",
        "subscribers_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/subscribers",
        "subscription_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/subscription",
        "commits_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/commits{/sha}",
        "git_commits_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/commits{/sha}",
        "comments_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/comments{/number}",
        "issue_comment_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/issues/comments{/number}",
        "contents_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/contents/{+path}",
        "compare_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/compare/{base}...{head}",
        "merges_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/merges",
        "archive_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/{archive_format}{/ref}",
        "downloads_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/downloads",
        "issues_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/issues{/number}",
        "pulls_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/pulls{/number}",
        "milestones_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/milestones{/number}",
        "notifications_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/notifications{?since,all,participating}",
        "labels_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/labels{/name}",
        "releases_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/releases{/id}",
        "deployments_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/deployments",
        "created_at": "2016-12-15T17:32:39Z",
        "updated_at": "2016-12-15T17:32:39Z",
        "pushed_at": "2016-12-15T17:32:41Z",
        "git_url": "git://github.com/Thamtham/gnampfelix-ci-tests.git",
        "ssh_url": "git@github.com:Thamtham/gnampfelix-ci-tests.git",
        "clone_url": "https://github.com/Thamtham/gnampfelix-ci-tests.git",
        "svn_url": "https://github.com/Thamtham/gnampfelix-ci-tests",
        "homepage": null,
        "size": 0,
        "stargazers_count": 0,
        "watchers_count": 0,
        "language": null,
        "has_issues": true,
        "has_downloads": true,
        "has_wiki": true,
        "has_pages": false,
        "forks_count": 0,
        "mirror_url": null,
        "open_issues_count": 0,
        "forks": 0,
        "open_issues": 0,
        "watchers": 0,
        "default_branch": "master"
      },
      "sender": {
        "login": "Thamtham",
        "id": 10077533,
        "avatar_url": "https://avatars.githubusercontent.com/u/10077533?v=3",
        "gravatar_id": "",
        "url": "https://api.github.com/users/Thamtham",
        "html_url": "https://github.com/Thamtham",
        "followers_url": "https://api.github.com/users/Thamtham/followers",
        "following_url": "https://api.github.com/users/Thamtham/following{/other_user}",
        "gists_url": "https://api.github.com/users/Thamtham/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/Thamtham/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/Thamtham/subscriptions",
        "organizations_url": "https://api.github.com/users/Thamtham/orgs",
        "repos_url": "https://api.github.com/users/Thamtham/repos",
        "events_url": "https://api.github.com/users/Thamtham/events{/privacy}",
        "received_events_url": "https://api.github.com/users/Thamtham/received_events",
        "type": "User",
        "site_admin": false
      }
    }
`

var notificationPingValidInvalidConfig = `
{
      "zen": "Encourage flow.",
      "hook_id": 11169236,
      "hook": {
        "type": "Repository",
        "id": 11169236,
        "name": "web",
        "active": true,
        "events": [
          "pull_request",
          "push"
        ],
        "config": {
          "content_type": "json",
          "insecure_ssl": "0",
          "url": "http://gnampfelix.de:8080/notification"
        },
        "updated_at": "2016-12-15T18:00:00Z",
        "created_at": "2016-12-15T18:00:00Z",
        "url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/hooks/11169236",
        "test_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/hooks/11169236/test",
        "ping_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/hooks/11169236/pings",
        "last_response": {
          "code": null,
          "status": "unused",
          "message": null
        }
      },
      "repository": {
        "id": 76581419,
        "name": "rolap",
        "full_name": "Thamtham/rolap",
        "owner": {
          "login": "Thamtham",
          "id": 10077533,
          "avatar_url": "https://avatars.githubusercontent.com/u/10077533?v=3",
          "gravatar_id": "",
          "url": "https://api.github.com/users/Thamtham",
          "html_url": "https://github.com/Thamtham",
          "followers_url": "https://api.github.com/users/Thamtham/followers",
          "following_url": "https://api.github.com/users/Thamtham/following{/other_user}",
          "gists_url": "https://api.github.com/users/Thamtham/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/Thamtham/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/Thamtham/subscriptions",
          "organizations_url": "https://api.github.com/users/Thamtham/orgs",
          "repos_url": "https://api.github.com/users/Thamtham/repos",
          "events_url": "https://api.github.com/users/Thamtham/events{/privacy}",
          "received_events_url": "https://api.github.com/users/Thamtham/received_events",
          "type": "User",
          "site_admin": false
        },
        "private": true,
        "html_url": "https://github.com/Thamtham/gnampfelix-ci-tests",
        "description": null,
        "fork": false,
        "url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests",
        "forks_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/forks",
        "keys_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/keys{/key_id}",
        "collaborators_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/collaborators{/collaborator}",
        "teams_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/teams",
        "hooks_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/hooks",
        "issue_events_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/issues/events{/number}",
        "events_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/events",
        "assignees_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/assignees{/user}",
        "branches_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/branches{/branch}",
        "tags_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/tags",
        "blobs_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/blobs{/sha}",
        "git_tags_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/tags{/sha}",
        "git_refs_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/refs{/sha}",
        "trees_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/trees{/sha}",
        "statuses_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/statuses/{sha}",
        "languages_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/languages",
        "stargazers_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/stargazers",
        "contributors_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/contributors",
        "subscribers_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/subscribers",
        "subscription_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/subscription",
        "commits_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/commits{/sha}",
        "git_commits_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/commits{/sha}",
        "comments_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/comments{/number}",
        "issue_comment_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/issues/comments{/number}",
        "contents_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/contents/{+path}",
        "compare_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/compare/{base}...{head}",
        "merges_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/merges",
        "archive_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/{archive_format}{/ref}",
        "downloads_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/downloads",
        "issues_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/issues{/number}",
        "pulls_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/pulls{/number}",
        "milestones_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/milestones{/number}",
        "notifications_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/notifications{?since,all,participating}",
        "labels_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/labels{/name}",
        "releases_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/releases{/id}",
        "deployments_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/deployments",
        "created_at": "2016-12-15T17:32:39Z",
        "updated_at": "2016-12-15T17:32:39Z",
        "pushed_at": "2016-12-15T17:32:41Z",
        "git_url": "git://github.com/Thamtham/gnampfelix-ci-tests.git",
        "ssh_url": "git@github.com:Thamtham/gnampfelix-ci-tests.git",
        "clone_url": "https://github.com/Thamtham/gnampfelix-ci-tests.git",
        "svn_url": "https://github.com/Thamtham/gnampfelix-ci-tests",
        "homepage": null,
        "size": 0,
        "stargazers_count": 0,
        "watchers_count": 0,
        "language": null,
        "has_issues": true,
        "has_downloads": true,
        "has_wiki": true,
        "has_pages": false,
        "forks_count": 0,
        "mirror_url": null,
        "open_issues_count": 0,
        "forks": 0,
        "open_issues": 0,
        "watchers": 0,
        "default_branch": "master"
      },
      "sender": {
        "login": "Thamtham",
        "id": 10077533,
        "avatar_url": "https://avatars.githubusercontent.com/u/10077533?v=3",
        "gravatar_id": "",
        "url": "https://api.github.com/users/Thamtham",
        "html_url": "https://github.com/Thamtham",
        "followers_url": "https://api.github.com/users/Thamtham/followers",
        "following_url": "https://api.github.com/users/Thamtham/following{/other_user}",
        "gists_url": "https://api.github.com/users/Thamtham/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/Thamtham/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/Thamtham/subscriptions",
        "organizations_url": "https://api.github.com/users/Thamtham/orgs",
        "repos_url": "https://api.github.com/users/Thamtham/repos",
        "events_url": "https://api.github.com/users/Thamtham/events{/privacy}",
        "received_events_url": "https://api.github.com/users/Thamtham/received_events",
        "type": "User",
        "site_admin": false
      }
    }
`

var notificationPingInvalid = `
{
      "zen": "Encourage flow.",
      "hook_id": 11169236,
      "hook": {
        "type": "Repository",
        "id": 11169236,
        "name": "web",
        "active": true,
        "events": [
          "pull_request",
          "push"
        ],
        "config": {
          "content_type": "json",
          "insecure_ssl": "0",
          "url": "http://gnampfelix.de:8080/notification"
        },
        "updated_at": "2016-12-15T18:00:00Z",
        "created_at": "2016-12-15T18:00:00Z",
        "url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/hooks/11169236",
        "test_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/hooks/11169236/test",
        "ping_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/hooks/11169236/pings",
        "last_response": {
          "code": null,
          "status": "unused",
          "message": null
        }
      },
        "private": true,
        "html_url": "https://github.com/Thamtham/gnampfelix-ci-tests",
        "description": null,
        "fork": false,
        "url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests",
        "forks_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/forks",
        "keys_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/keys{/key_id}",
        "collaborators_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/collaborators{/collaborator}",
        "teams_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/teams",
        "hooks_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/hooks",
        "issue_events_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/issues/events{/number}",
        "events_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/events",
        "assignees_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/assignees{/user}",
        "branches_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/branches{/branch}",
        "tags_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/tags",
        "blobs_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/blobs{/sha}",
        "git_tags_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/tags{/sha}",
        "git_refs_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/refs{/sha}",
        "trees_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/trees{/sha}",
        "statuses_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/statuses/{sha}",
        "languages_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/languages",
        "stargazers_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/stargazers",
        "contributors_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/contributors",
        "subscribers_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/subscribers",
        "subscription_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/subscription",
        "commits_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/commits{/sha}",
        "git_commits_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/git/commits{/sha}",
        "comments_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/comments{/number}",
        "issue_comment_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/issues/comments{/number}",
        "contents_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/contents/{+path}",
        "compare_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/compare/{base}...{head}",
        "merges_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/merges",
        "archive_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/{archive_format}{/ref}",
        "downloads_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/downloads",
        "issues_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/issues{/number}",
        "pulls_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/pulls{/number}",
        "milestones_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/milestones{/number}",
        "notifications_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/notifications{?since,all,participating}",
        "labels_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/labels{/name}",
        "releases_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/releases{/id}",
        "deployments_url": "https://api.github.com/repos/Thamtham/gnampfelix-ci-tests/deployments",
        "created_at": "2016-12-15T17:32:39Z",
        "updated_at": "2016-12-15T17:32:39Z",
        "pushed_at": "2016-12-15T17:32:41Z",
        "git_url": "git://github.com/Thamtham/gnampfelix-ci-tests.git",
        "ssh_url": "git@github.com:Thamtham/gnampfelix-ci-tests.git",
        "clone_url": "https://github.com/Thamtham/gnampfelix-ci-tests.git",
        "svn_url": "https://github.com/Thamtham/gnampfelix-ci-tests",
        "homepage": null,
        "size": 0,
        "stargazers_count": 0,
        "watchers_count": 0,
        "language": null,
        "has_issues": true,
        "has_downloads": true,
        "has_wiki": true,
        "has_pages": false,
        "forks_count": 0,
        "mirror_url": null,
        "open_issues_count": 0,
        "forks": 0,
        "open_issues": 0,
        "watchers": 0,
        "default_branch": "master"
      },
      "sender": {
        "login": "Thamtham",
        "id": 10077533,
        "avatar_url": "https://avatars.githubusercontent.com/u/10077533?v=3",
        "gravatar_id": "",
        "url": "https://api.github.com/users/Thamtham",
        "html_url": "https://github.com/Thamtham",
        "followers_url": "https://api.github.com/users/Thamtham/followers",
        "following_url": "https://api.github.com/users/Thamtham/following{/other_user}",
        "gists_url": "https://api.github.com/users/Thamtham/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/Thamtham/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/Thamtham/subscriptions",
        "organizations_url": "https://api.github.com/users/Thamtham/orgs",
        "repos_url": "https://api.github.com/users/Thamtham/repos",
        "events_url": "https://api.github.com/users/Thamtham/events{/privacy}",
        "received_events_url": "https://api.github.com/users/Thamtham/received_events",
        "type": "User",
        "site_admin": false
      }
    }
`
var validConfigFile = `
{
    "push":[],
    "pull_request":[]
}
`
