package test_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// "github.com/rg-km/final-project-engineering-12/backend/test"
)

var _ = Describe("Test", func() {
	Describe("Add", func() {

		It("adds two numbers", func() {
			sum := 5
			Expect(sum).To(Equal(5))
		})
	})
})
