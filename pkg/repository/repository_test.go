package repository_test

import (
	. "github.com/gnampfelix/gnampfelix-ci/pkg/repository"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository", func() {
	var repo Repository
	It("should clone the repository", func() {
		repo = New()
		repo.SetUrl("https://github.com/gnampfelix/pub")
		repo.SetName("pub")
		repo.SetId("abcdefg")
		repo.SetRef("refs/heads/master")
		err := repo.EnsureLocalExistence()
		Expect(err).Should(Succeed())
		Expect("abcdefg").Should(BeADirectory())
	})
	It("should delete the repository", func() {
		repo.Delete()
		Expect("abcdefg").Should(Not(BeADirectory()))
	})
})
