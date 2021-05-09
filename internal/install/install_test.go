package install_test

import (
	"StackDB/internal/install"
	"StackDB/internal/utils"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Install", func() {
	BeforeEach(func() {
		os.Chdir("../../")
		utils.GetEnv()
	})
	AfterEach(func() {
		os.RemoveAll("./stackdb")
	})
	Describe("Installing StackDB", func() {
		Context("When completed successfully", func() {
			It("should return nil", func() {
				Expect(install.Intall()).To(BeNil())
			})
		})
	})
})
