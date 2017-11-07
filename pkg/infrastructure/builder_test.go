package infrastructure_test

import (
	"github.com/docker/docker/client"
	. "github.com/gnampfelix/gnampfelix-ci/pkg/infrastructure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
)

//	I am not sure how to structure those test. Need to rework!
var _ = Describe("Builder - require Docker", func() {
	var env BuildEnvironment
	var err error

	docker, _ := client.NewEnvClient()

	It("should return the default builder", func() {
		env, err = GetEnvironment("default", "")
		Expect(err).Should(Succeed())
		Expect(env).Should(Not(BeNil()))
	})

	Context("docker", func() {
		BeforeEach(func() {
			err = os.Mkdir("testId", os.ModePerm)
			Expect(err).Should(Succeed())
			file, _ := os.Create("testId/pre.sh")
			file.Write([]byte(pre))
			file.Sync()
			file.Close()
		})
		AfterEach(func() {
			os.RemoveAll("testId")
		})

		It("should run the complete lifecycle", func() {
			err = env.Create("testId")
			Expect(err).Should(Succeed())

			ctx := context.Background()
			status, err := docker.ContainerInspect(ctx, "testId")
			Expect(err).Should(Succeed())
			Expect(status.State.Status).Should(Equal("created"))

			err = env.StartBuild()
			Expect(err).Should(Succeed())

			output := env.OutputStream()
			outputAll, err := ioutil.ReadAll(output)

			Expect(err).Should(Succeed())
			Expect(outputAll).Should(Equal([]byte(expectedResult)))

			buildResult := env.Wait()
			Expect(buildResult).Should(Equal(BuildResultTestError))
			env.Destroy()
		})
	})
})

/*
	It's time for: Fun Facts! TTYs convert \n into \r\n, so we have to input the \r\n explicitly!
*/
var expectedResult string = "HI\r\n/bin/bash: /build/test.sh: No such file or directory\r\n"

var pre string = `
#!/bin/bash
echo "HI"
`
