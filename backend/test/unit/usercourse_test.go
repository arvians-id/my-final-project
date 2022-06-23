package unit_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-12/backend/config"
	"github.com/rg-km/final-project-engineering-12/backend/test/setup"
)

var _ = Describe("User Course API", func() {
	var server *gin.Engine

	BeforeEach(func() {
		configuration := config.New("../../.env.test")

		_, err := setup.SuiteSetup(configuration)
		if err != nil {
			panic(err)
		}

		router := setup.ModuleSetup(configuration)
		server = router
	})

	AfterEach(func() {
		configuration := config.New("../../.env.test")
		db, err := setup.SuiteSetup(configuration)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = setup.TearDownTest(db)
		if err != nil {
			panic(err)
		}
	})

	Describe("Create User Course", func() {
		When("Data is empty", func() {
			It("should return All User Course", func() {
				// Create User Course 1
				requestBody := strings.NewReader(`{"user_id": 1,"course_id": 1}`)
				request := httptest.NewRequest(http.MethodPost, "/api/usercourses", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Create User Course 2
				requestBody = strings.NewReader(`{"user_id": 1,"course_id": 2}`)
				request = httptest.NewRequest(http.MethodPost, "/api/usercourses", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Create User Course 3
				requestBody = strings.NewReader(`{"user_id": 2,"course_id": 1}`)
				request = httptest.NewRequest(http.MethodPost, "/api/usercourses", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// find all user courses
				request = httptest.NewRequest(http.MethodGet, "/api/usercourses", nil)
				request.Header.Add("Content-Type", "application/json")

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				Body, err := ioutil.ReadAll(writer.Body)
				if err != nil {
					panic(err)
				}
				var responseBody map[string]interface{}
				_ = json.Unmarshal(Body, &responseBody)

				usercourse := responseBody["data"].([]interface{})
				usercourse1 := usercourse[0].(map[string]interface{})
				usercourse2 := usercourse[1].(map[string]interface{})
				usercourse3 := usercourse[2].(map[string]interface{})

				Expect(int(responseBody["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody["status"]).To(Equal("OK"))

				Expect(usercourse1["user_id"]).To(Equal(1))
				Expect(usercourse1["course_id"]).To(Equal(1))

				Expect(usercourse2["user_id"]).To(Equal(1))
				Expect(usercourse2["course_id"]).To(Equal(2))

				Expect(usercourse3["user_id"]).To(Equal(2))
				Expect(usercourse3["course_id"]).To(Equal(1))
			})
		})
	})
})
