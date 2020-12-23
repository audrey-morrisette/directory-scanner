package directoryscanner_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/audrey-morrisette/directoryscanner"
)

var _ = Describe("Directoryscan", func() {
	Describe("Scanning a folder with 16 known findings", func() {
		Context("less than 100 MB in size", func() {
			It("should return 16 findings", func() {
				Expect(len(directoryscanner.Dig("../"))).To(Equal(16))
			})
		})
	})
})
