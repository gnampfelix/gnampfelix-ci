package repository_test

import (
	. "github.com/gnampfelix/gnampfelix-ci/pkg/repository"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("RepoConfig", func() {
	var r RepoConfig
	var a Action
	var t Task
	It("should parse the config file", func() {
		var err error
		file, _ := os.Create("test1.json")
		file.Write([]byte(test1))
		file.Sync()
		file.Close()
		r, err = Parse("test1.json")
		Expect(err).Should(Succeed())
		Expect(r).Should(Not(BeNil()))
		os.Remove("test1.json")
	})

	It("should return an action", func() {
		var err error
		a, err = r.GetAction("testBranch")
		Expect(err).Should(Succeed())
		Expect(a).Should(Not(BeNil()))
	})

	It("should return a task", func() {
		t = a.GetTask()
		Expect(t).Should(Not(BeNil()))
	})

	It("should return \"default\" environment", func() {
		env := a.GetEnvironment()
		Expect(env).Should(Not(BeNil()))
	})

	It("should return pre-script", func() {
		pre := t.GetPre()
		Expect(pre).Should(Equal("abc.sh"))
	})

	It("should return test-script", func() {
		test := t.GetTest()
		Expect(test).Should(Equal("def.sh"))
	})

	It("should return post-script", func() {
		post := t.GetPost()
		Expect(post).Should(Equal("ghi.sh"))
	})

	It("should reject the file (missing tasks-section)", func() {
		var err error
		file, _ := os.Create("test2.json")
		file.Write([]byte(test2))
		file.Sync()
		file.Close()
		r, err = Parse("test2.json")
		Expect(err).Should(HaveOccurred())
		Expect(r).Should(BeNil())
		os.Remove("test2.json")
	})

	It("should reject the file (invalid task reference)", func() {
		var err error
		file, _ := os.Create("test3.json")
		file.Write([]byte(test3))
		file.Sync()
		file.Close()
		r, err = Parse("test3.json")
		Expect(err).Should(HaveOccurred())
		Expect(r).Should(BeNil())
		os.Remove("test3.json")
	})

	It("should reject the file (invalid task declaration)", func() {
		var err error
		file, _ := os.Create("test4.json")
		file.Write([]byte(test4))
		file.Sync()
		file.Close()
		r, err = Parse("test4.json")
		Expect(err).Should(HaveOccurred())
		Expect(r).Should(BeNil())
		os.Remove("test4.json")
	})
})

var test1 string = `
{
    "tasks": {
        "task1": {
            "pre": "abc.sh",
            "test": "def.sh",
            "post": "ghi.sh"
        }
    },
    "actions": {
        "{{ all }}": {
            "env": "default",
            "task": "task1"
        }
    }
}
`

var test2 string = `
{
    "actions": {
        "{{ all }}": {
            "env": "default",
            "task": "task1"
        }
    }
}}`

var test3 string = `
{
    "tasks": {
        "task1": {
            "pre": "abc.sh",
            "test": "def.sh",
            "post": "ghi.sh"
        }
    },
    "actions": {
        "{{ all }}": {
            "env": "default",
            "task": "task2"
        }
    }
}
`

var test4 string = `
{
    "tasks": {
        "task1": {
            "pre": "abc.sh",
            "test": "def.sh"
        }
    },
    "actions": {
        "{{ all }}": {
            "env": "default",
            "task": "task2"
        }
    }
}
`
