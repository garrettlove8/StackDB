package collection_test

import (
	"StackDB/internal/collection"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Collections", func() {
	// BeforeEach(func() {
	// 	os.Chdir("../../")
	// 	utils.GetEnv()
	// })
	// AfterEach(func() {
	// 	os.RemoveAll("./sdb")
	// })
	Describe("Create New Collection", func() {
		Context("When created successfully", func() {
			It("should have non-nil values for the appropriate properties", func() {
				newCol := collection.NewCollection()
				Expect(newCol.Uuid).NotTo(BeNil())
				Expect(newCol.CTime).NotTo(BeNil())
				Expect(newCol.UTime).NotTo(BeNil())
			})
		})
	})
	Describe("Create New Data", func() {
		Context("When created successfully", func() {
			It("should have non-nil values for the appropriate properties", func() {
				newData := collection.NewData()
				Expect(newData.Uuid).NotTo(BeNil())
				Expect(newData.CTime).NotTo(BeNil())
				Expect(newData.UTime).NotTo(BeNil())
			})
		})
	})
})
