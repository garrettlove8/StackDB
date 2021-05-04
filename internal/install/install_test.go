package install_test

import (
	"StackDB/internal/database"
	"StackDB/internal/install"
	"os"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Install", func() {
	AfterEach(func() {
		os.RemoveAll("./stackdb")
	})
	Describe("Installing StackDB", func() {
		Context("When completed successfully", func() {
			It("should return nil", func() {
				systemDb := database.Database{
					Id:   uuid.New().String(),
					Name: "system",
					Type: "keyValue",
				}
				Expect(install.Intall(&systemDb)).To(BeNil())
			})
		})
	})
})
