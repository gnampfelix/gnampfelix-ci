package infrastructure

import (
	// "github.com/docker/docker/client"
	"io"
)

type BuildEnvironment interface {
	Create(buildId string) error
	StartBuild() error
	OutputStream() io.ReadCloser
	Wait() //BuildResult
	Destroy()
}

//	Returns a BuildEnvironment that is specified with identifier. If identifier
//	leads to a complex environment, path is searched for more details.
func GetEnvironment(identifier, path string) (BuildEnvironment, error) {
	return nil, nil
}
