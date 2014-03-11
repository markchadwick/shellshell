package shellshell

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Builder", func() {
	var b *Builder

	BeforeEach(func() {
		b = New()
	})

	It("should have an empty default prompt", func() {
		Expect(b.Prompt()).To(Equal(""))
	})

	It("should set a prompt", func() {
		b.SetPrompt(func() string { return "pants> " })
		Expect(b.Prompt()).To(Equal("pants> "))
	})
})
