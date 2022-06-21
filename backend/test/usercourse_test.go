package test_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-12/backend/controller"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

var _ = Describe("Test", func() {
	var usercoursecreate model.CreateUserCourseRequest
	var usercoursecontroller controller.UserCourseController
	var usercourse entity.UserCourse
	var ctx *gin.Context

	usercoursecreate = model.CreateUserCourseRequest{
		UserId:   5,
		CourseId: 3,
	}

	usercourse = entity.UserCourse{
		UserId:   5,
		CourseId: 3,
	}

	jsonByte, _ := json.Marshal(usercoursecreate)
	ctx = &gin.Context{}
	ctx.Request = &http.Request{
		Body: ioutil.NopCloser(bytes.NewBuffer(jsonByte)),
	}

	Describe("usercase_test", func() {

		When("usercourse create", func() {
			It("should create a new usercourse", func() {
				responses, err := usercoursecontroller.UserCourseService.Create(ctx, usercoursecreate)
				Expect(err).To(BeNil())
				Expect(responses).To(Equal(usercourse))
				Expect(usercourse.UserId).To(Equal(5))
				Expect(usercourse.CourseId).To(Equal(3))
			})
		})

		When("usercourse findbyuserid", func() {
			It("Should get a usercourse by userid", func() {
				responses, err := usercoursecontroller.UserCourseService.FindByUserId(ctx, utils.ToString(usercoursecreate.UserId))
				Expect(err).To(BeNil())
				Expect(responses).To(Equal(usercourse))
				Expect(usercourse.UserId).To(Equal(5))
				Expect(usercourse.CourseId).To(Equal(3))
			})
		})

		When("usercourse list", func() {
			It("Should get a list of usercourse", func() {
				responses, err := usercoursecontroller.UserCourseService.FindAll(ctx)
				Expect(err).To(BeNil())
				Expect(responses).To(Equal(usercourse))
				Expect(usercourse.UserId).To(Equal(5))
				Expect(usercourse.CourseId).To(Equal(3))
			})
		})

		When("usercourse delete", func() {
			It("Should delete a usercourse", func() {
				err := usercoursecontroller.UserCourseService.Delete(ctx, usercoursecreate.UserId, usercoursecreate.CourseId)
				Expect(err).To(BeNil())
			})
		})
	})
})
