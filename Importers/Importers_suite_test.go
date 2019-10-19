package Importers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestImporters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Importers Suite")
}
