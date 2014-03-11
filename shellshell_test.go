package shellshell

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shellshell Suite")
}

var _ = Describe("Completer", func() {
	var c *Completer

	BeforeEach(func() {
		c = NewCompleter()
	})

	It("should return no results by default", func() {
		Expect(c.Complete("")).To(BeEmpty())
		Expect(c.Complete("anything")).To(BeEmpty())
		Expect(c.Complete("*")).To(BeEmpty())
	})

	It("should complete a simple prefix", func() {
		c.Prefix("hello", func(l string) []string {
			return []string{"there"}
		})
		Expect(c.Complete("hel")).ToNot(BeEmpty())
		Expect(c.Complete("hel")).To(HaveLen(1))
		Expect(c.Complete("hel")[0]).To(Equal("there"))
	})

	It("should allow multiple completion value", func() {
		c.Prefix("dogs", func(l string) []string {
			return []string{"dig", "dug"}
		})
		res := c.Complete("dog")
		Expect(res).To(HaveLen(2))
		Expect(res[0]).To(Equal("dig"))
		Expect(res[1]).To(Equal("dug"))
	})

	It("should register multiple handlers", func() {
		c.Prefix("dogs", func(l string) []string {
			return []string{"was dogs"}
		})
		c.Prefix("doggies", func(l string) []string {
			return []string{"was doggies"}
		})
		c.Prefix("cats", func(l string) []string {
			return []string{"was cats"}
		})
		res := c.Complete("dog")
		Expect(res).To(HaveLen(2))
		Expect(res[0]).To(Equal("was dogs"))
		Expect(res[1]).To(Equal("was doggies"))
	})
})
