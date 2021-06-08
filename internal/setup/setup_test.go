package setup_test

import (
	"StackDB/internal/setup"
	"StackDB/internal/utils"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Setup", func() {
	BeforeSuite(func() {
		os.Chdir("../../")
		utils.GetEnv()
	})
	AfterEach(func() {
		os.RemoveAll("./sdb")
	})
	Describe("Setup StackDB", func() {
		Context("When completed successfully", func() {
			It("should return nil", func() {
				Expect(setup.Setup()).To(BeNil())
			})
		})
	})
	Describe("Check setup", func() {
		Context("When setup has not yet been run", func() {
			It("should return false", func() {
				Expect(setup.CheckSetup()).To(Equal(false))
			})
		})
	})
})
