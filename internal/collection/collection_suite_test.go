package collection_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInstall(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Collection Suite")
}
