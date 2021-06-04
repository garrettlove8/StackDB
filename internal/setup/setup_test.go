package setup_test

import (
	"StackDB/internal/setup"
	"StackDB/internal/utils"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Setup", func() {
	BeforeEach(func() {
		os.Chdir("../../")
		utils.GetEnv()
	})
	AfterEach(func() {
		os.RemoveAll("./sdb")
	})
	Describe("Setup StackDB", func() {
		Context("When completed successfully", func() {
			XIt("should return nil", func() {
				Expect(setup.Setup()).To(BeNil())
			})
		})
	})
})
