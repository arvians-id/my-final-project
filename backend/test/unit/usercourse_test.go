package unit_test

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-12/backend/config"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/test/setup"
)

var _ = Describe("User Course API", func() {
	var (
		server *gin.Engine
		token  string
		ok     bool
	)

	BeforeEach(func() {
		configuration := config.New("../../.env.test")

		_, err := setup.SuiteSetup(configuration)
		if err != nil {
			panic(err)
		}

		router := setup.ModuleSetup(configuration)
		server = router

		var user = model.UserRegisterResponse{

			Name:           "akuntest",
			Username:       "akuntest",
			Email:          "akuntest@gmail.com",
			Password:       "123456ll",
			Role:           1,
			Phone:          "085156789011",
			Gender:         1,
			DisabilityType: 1,
			Birthdate:      "2002-04-01",
		}

		login := model.GetUserLogin{

			Email:    "akuntest@gmail.com",
			Password: "123456ll",
		}

		// Register User
		userData, _ := json.Marshal(user)
		requestBody := strings.NewReader(string(userData))
		request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
		request.Header.Add("Content-Type", "application/json")

		writer := httptest.NewRecorder()
		server.ServeHTTP(writer, request)

		//Login User
		userLogin, _ := json.Marshal(login)
		requestBodyLogin := strings.NewReader(string(userLogin))
		requestLogin := httptest.NewRequest(http.MethodPost, "/api/users/login", requestBodyLogin)
		request.Header.Add("Content-Type", "application/json")

		writerLogin := httptest.NewRecorder()
		server.ServeHTTP(writerLogin, requestLogin)

		responseLogin := writerLogin.Result()

		bodyLogin, _ := io.ReadAll(responseLogin.Body)
		var responseBodyLogin map[string]interface{}
		_ = json.Unmarshal(bodyLogin, &responseBodyLogin)

		token, ok = responseBodyLogin["token"].(string)
		if !ok {
			panic("Can't get token")
		} else {
			log.Println("Token: ", token)
		}
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

	Describe("Create User Course and List User Course", func() {
		When("Data is empty", func() {
			It("should return All User Course", func() {
				// Create User Course 1
				requestBody := strings.NewReader(`{"user_id": 10,"course_id": 1}`)
				request := httptest.NewRequest(http.MethodPost, "/api/usercourse", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Create User Course 2
				requestBody = strings.NewReader(`{"user_id": 10,"course_id": 2}`)
				request = httptest.NewRequest(http.MethodPost, "/api/usercourse", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Create User Course 3
				requestBody = strings.NewReader(`{"user_id": 11,"course_id": 1}`)
				request = httptest.NewRequest(http.MethodPost, "/api/usercourse", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// find all user courses
				request = httptest.NewRequest(http.MethodGet, "/api/usercourse", nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				Body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(Body, &responseBody)

				usercourse := responseBody["data"].([]interface{})
				usercourse1 := usercourse[0].(map[string]interface{})
				usercourse2 := usercourse[1].(map[string]interface{})
				usercourse3 := usercourse[2].(map[string]interface{})

				Expect(int(responseBody["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody["status"]).To(Equal("OK"))

				Expect(usercourse1["user_id"]).To(Equal(10))
				Expect(usercourse1["course_id"]).To(Equal(1))

				Expect(usercourse2["user_id"]).To(Equal(10))
				Expect(usercourse2["course_id"]).To(Equal(2))

				Expect(usercourse3["user_id"]).To(Equal(11))
				Expect(usercourse3["course_id"]).To(Equal(1))
			})
		})
	})
	Describe("Get User Course By UserId and CourseId", func() {
		When("Data is empty", func() {
			It("should return User Course", func() {
				// Create User Course 1
				requestBody := strings.NewReader(`{"user_id": 10,"course_id": 1}`)
				request := httptest.NewRequest(http.MethodPost, "/api/usercourse", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)
				// Get User Course By UserId and CourseId
				request = httptest.NewRequest(http.MethodGet, "/api/usercourse/10/1", nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				Body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(Body, &responseBody)

				usercourse := responseBody["data"].(map[string]interface{})

				Expect(usercourse["user_id"]).To(Equal(10))
				Expect(usercourse["course_id"]).To(Equal(1))
			})
		})
	})
	Describe("Delete User Course By UserId and CourseId", func() {
		When("Data is empty", func() {
			It("should return User Course", func() {
				// Create User Course 1
				requestBody := strings.NewReader(`{"user_id": 10,"course_id": 1}`)
				request := httptest.NewRequest(http.MethodPost, "/api/usercourse", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)
				// Delete User Course By UserId and CourseId
				request = httptest.NewRequest(http.MethodDelete, "/api/usercourse/10/1", nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				Body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(Body, &responseBody)

				Expect(int(responseBody["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody["status"]).To(Equal("OK"))
			})
		})
	})
})
