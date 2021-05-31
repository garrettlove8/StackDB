package shell_test

import (
	"StackDB/internal/setup"
	"StackDB/internal/shell"
	"StackDB/internal/utils"
	"errors"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shell", func() {
	BeforeSuite(func() {
		os.Chdir("../../")
		utils.GetEnv()
	})
	AfterEach(func() {
		os.RemoveAll("./sdb")
	})
	Describe("Starting the shell", func() {
		Context("When starting successfully", func() {
			XIt("should return nil", func() {
				setup.Setup()
				Expect(shell.Start()).To(BeNil())
			})
		})
		Context("When starting unsuccessfully", func() {
			XIt("should inform the caller by returning an error", func() {
				err := errors.New("unable to load database stackdb")
				Expect(shell.Start()).To(Equal(err))
			})
		})
	})
})
