package integration

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMainApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Testing")
}
