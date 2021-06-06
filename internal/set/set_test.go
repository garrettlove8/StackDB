package set_test

import (
	"StackDB/internal/set"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sets", func() {
	// BeforeEach(func() {
	// 	os.Chdir("../../")
	// 	utils.GetEnv()
	// })
	// AfterEach(func() {
	// 	os.RemoveAll("./sdb")
	// })
	Describe("Successfully create new Set", func() {
		Context("When using default values", func() {
			It("should have correct value / non-nil propety values", func() {
				newCol, _ := set.NewSet("setName")

				Expect(newCol.Name).To(Equal("setName"))
				Expect(newCol.Uuid).NotTo(BeNil())
				Expect(newCol.CTime).NotTo(BeNil())
				Expect(newCol.UTime).NotTo(BeNil())
			})
		})
		Context("When using custom values", func() {
			It("should have used the custom values", func() {
				newCol, _ := set.NewSet("setName", "setUuid", "setLocation")
				Expect(newCol.Name).To(Equal("setName"))
				Expect(newCol.Uuid).To(Equal("setUuid"))
				Expect(newCol.Location).To(Equal("setLocation"))
			})
		})
	})
	Describe("Unsuccessfully create new Set", func() {
		Context("When not supplying a name", func() {
			It("should return an error indicating that no name was provided", func() {
				_, err := set.NewSet()

				wantedErr := errors.New("no name provided for new Set.")
				Expect(err).To(Equal(wantedErr))
			})
		})
	})
})
