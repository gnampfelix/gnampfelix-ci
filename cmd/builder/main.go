package main

import (
	"flag"
	"fmt"
	"github.com/gnampfelix/pub"
    "github.com/gnampfelix/gnampfelix-ci/pkg/common"
	git "gopkg.in/src-d/go-git.v4"
	gitHttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
    "gopkg.in/src-d/go-git.v4/plumbing"
	"io/ioutil"
	"os"
    "sync"
    "os/exec"
    "errors"
)

var (
	gitUser     string
	gitToken    string
	gitUrl      string
	gitRepoName string
    gitRef      string
    preTest     *os.File
)

func main() {
    buildId, serverUrl, err := parseVariables()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    connection, err := prepareConnection(buildId, serverUrl)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	defer connection.Close()

	err = getGitData(connection)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

    waitgroup := sync.WaitGroup{}
    waitgroup.Add(2)
    errChan := make(chan error, 2)
    go cloneAndCheckout(&waitgroup, errChan)
    go getTestScripts(&waitgroup, errChan, connection)
    waitgroup.Wait()

    if err = evalErrChan(errChan); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    cmd := exec.Command("bash", "preTest.sh")
    connection.StartStream()
    cmd.Stdout = connection
    err = cmd.Run()
    if err != nil {
        connection.StopStream()
        connection.SendMessage(common.ErrorMessage())
        return
    }

    cmd := exec.Command("bash", "test.sh")
    cmd.Stdout(connection)
    err = cmd.Run()
    if err != nil {
        connection.StopStream()
        connection.SendMessage(common.FailureMessage())
        return
    }

    connection.StopStream()
    connection.SendMessage(common.SuccessMessage())
}

func evalErrChan(errChan chan error) error {
    if err := <- errChan; err != nil {
        return err
    }
    err := <- errChan
    return err
}

func parseVariables() (string, string, error) {
    var buildId string
    var serverUrl string
    flag.StringVar(&buildId, "build", "", "The build id of the current build run.")
    flag.StringVar(&serverUrl, "master", "", "The url of the master server.")
    flag.Parse()

    if buildId == "" || serverUrl == "" {
        return "", "", errors.New("missing build or master information")
    }
    return buildId, serverUrl, nil
}

func prepareConnection(connId, remote string) (pub.Connection, error) {
    messenger := pub.NewMessenger()
	connection, err := messenger.TalkTo(remote)
	if err != nil {
		return nil, err
	}

    err = connection.SendString(connId)
    if err != nil {
        connection.Close()
    }
    return connection, err
}

func getGitData(connection pub.Connection) error {
	messageUser, err := connection.ReceiveMessageWithTag("gitUser")
	if err != nil {
		return err
	}
	gitUser, err = messageToString(messageUser)
    if err != nil {
        return err
    }

	messageToken, err := connection.ReceiveMessageWithTag("gitToken")
	if err != nil {
		return err
	}
	gitToken, err = messageToString(messageToken)
    if err != nil {
        return err
    }

	messageUrl, err := connection.ReceiveMessageWithTag("gitUrl")
	if err != nil {
		return err
	}
    gitUrl, err = messageToString(messageUrl)

    messageRef, err := connection.ReceiveMessageWithTag("gitRef")
    if err != nil {
        return err
    }
    gitRef, err = messageToString(messageRef)

	messageRepo, err := connection.ReceiveMessageWithTag("gitRepoName")
	if err != nil {
		return err
	}
	gitRepoName, err = messageToString(messageRepo)
	return err
}

func cloneAndCheckout(wg *sync.WaitGroup, errChan chan error) {
    defer wg.Done()
    auth := gitHttp.NewBasicAuth(gitUser, gitToken)
    repository, err := git.PlainClone(gitRepoName, false, &git.CloneOptions{
        URL: gitUrl,
        Auth: auth,
    })
    if err != nil {
        errChan <- err
        return
    }
    wt, err := repository.Worktree()
    if err != nil {
        errChan <- err
        return
    }

    err = wt.Checkout(&git.CheckoutOptions{Branch: plumbing.ReferenceName(gitRef)})
    errChan <- err
    return
}

func getTestScripts(wg *sync.WaitGroup, errChan chan error, connection pub.Connection) error {
    defer wg.Done()

    preTest, err = connection.ReceiveFile("preTest.sh")
    if err != nil {
        errChan <- err
        return
    }

    test, err = connection.ReceiveFile("test.sh")
    errChan <- err
    return
}

func messageToString(message pub.Message) (string, error) {
    messageContent, err := ioutil.ReadAll(message)
    if err != nil {
        return "", err
    }
    result := string(messageContent[:len(messageContent)-1])
    return result, nil
}
