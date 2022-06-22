package integration

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestMainApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Testing")
}
