package infrastructure

import (
	"bytes"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"os"
)

var docker *client.Client

func init() {
	var err error
	docker, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}
}

type BuildEnvironment interface {
	Create(buildId string) error
	StartBuild() error
	OutputStream() io.ReadCloser
	Wait() BuildResult
	Destroy()
}

//	Returns a BuildEnvironment that is specified with identifier. If identifier
//	leads to a complex environment, path is searched for more details.
func GetEnvironment(identifier, path string) (BuildEnvironment, error) {
	switch identifier {
	case "default":
		return &defaultBuilder{}, nil
	}
	return nil, errors.New("environment \"" + identifier + "\" is not a valid environment")
}

type defaultBuilder struct {
	buildId     string
	containerId string
}

func (d *defaultBuilder) Create(buildId string) error {
	d.buildId = buildId
	ctx := context.Background()

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	binds := []string{
		pwd + "/" + buildId + ":/build",
	}

	cont, err := docker.ContainerCreate(ctx, &container.Config{
		Image:        "gnampfelix/ci-builder",
		AttachStdout: true,
		Tty:          true,
	}, &container.HostConfig{
		Binds: binds,
	}, nil, buildId)
	if err != nil {
		return err
	}
	d.containerId = cont.ID
	return nil
}

func (d *defaultBuilder) StartBuild() error {
	ctx := context.Background()
	err := docker.ContainerStart(ctx, d.containerId, types.ContainerStartOptions{})
	return err
}

func (d *defaultBuilder) OutputStream() io.ReadCloser {
	ctx := context.Background()
	result, err := docker.ContainerLogs(ctx, d.containerId, types.ContainerLogsOptions{
		Follow:     true,
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		tmp := bytes.NewBufferString(err.Error())
		result = ioutil.NopCloser(tmp)
	}
	return result
}

func (d *defaultBuilder) Wait() BuildResult {
	ctx := context.Background()
	resultChan, errChan := docker.ContainerWait(ctx, d.containerId, container.WaitConditionNotRunning)
	select {
	case resultBody := <-resultChan:
		return IntToBuildResult(resultBody.StatusCode)
	case <-errChan:
		return BuildResultError
	}
}

func (d *defaultBuilder) Destroy() {
	ctx := context.Background()
	_ = docker.ContainerRemove(ctx, d.containerId, types.ContainerRemoveOptions{})
}
