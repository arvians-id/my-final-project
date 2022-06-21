package test_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	// "github.com/rg-km/final-project-engineering-12/backend/test"
)

var _ = Describe("Test", func() {
	var dummy entity.UserCourse
	var uscourse model.CreateUserCourseRequest

	dummy = entity.UserCourse{
		UserId:   5,
		CourseId: 3,
	}

	


	fmt.Println(dummy)
	fmt.Println(uscourse)






	Describe("Add", func() {

		It("adds two numbers", func() {
			sum := 5
			Expect(sum).To(Equal(5))
		})
	})

})
