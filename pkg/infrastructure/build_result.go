package infrastructure

type BuildResult int64

const (
	BuildResultSuccess   BuildResult = 0
	BuildResultPreError  BuildResult = 1
	BuildResultTestError BuildResult = 2
	BuildResultPostError BuildResult = 3
	BuildResultError                 = 4
)

func IntToBuildResult(input int64) BuildResult {
	if input >= 0 && input <= 4 {
		return BuildResult(input)
	} else {
		return BuildResultError
	}
}
