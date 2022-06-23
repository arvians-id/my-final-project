package unit_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-12/backend/config"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"github.com/rg-km/final-project-engineering-12/backend/test/setup"
)

var _ = Describe("User Course API", func() {
	var (
		server   *gin.Engine
		tokenJWT string
		ctx      *gin.Context
	)

	BeforeEach(func() {
		configuration := config.New("../../.env.test")

		db, err := setup.SuiteSetup(configuration)
		if err != nil {
			panic(err)
		}

		router := setup.ModuleSetup(configuration)
		database := db
		server = router

		// User Authentication
		// Create user
		tx, _ := database.Begin()
		userRepository := repository.NewUserRepository()
		userService := service.NewUserService(&userRepository, database)

		_, err = userService.RegisterUser(ctx, model.UserRegisterResponse{
			Id:                9,
			Name:              "test",
			Username:          "test",
			Email:             "test123@gmail.com",
			Password:          "Test123",
			Role:              1,
			Phone:             "081234567890",
			Gender:            1,
			DisabilityType:    1,
			Birthdate:         "2001-01-01",
			EmailVerification: time.Now(),
			Created_at:        time.Now(),
			Updated_at:        time.Now(),
		})
		if err != nil {
			panic(err)
		}
		err = tx.Commit()
		if err != nil {
			panic(err)
		}

		// Login user
		requestBody := strings.NewReader(`{"email": "test123@gmail.com","password": "Test123"}`)
		request := httptest.NewRequest(http.MethodPost, "/api/users/login", requestBody)
		request.Header.Add("Content-Type", "application/json")

		writer := httptest.NewRecorder()
		server.ServeHTTP(writer, request)

		response := writer.Result()

		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]string
		_ = json.Unmarshal(body, &responseBody)

		tokenJWT = responseBody["token"]
		if tokenJWT == "" {
			panic("Token is empty")
		} else {
			fmt.Println("Token: ", tokenJWT)
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

	Describe("Create User Course", func() {
		When("Data is empty", func() {
			It("should return All User Course", func() {
				// Create User Course 1
				requestBody := strings.NewReader(`{"user_id": 10,"course_id": 1}`)
				request := httptest.NewRequest(http.MethodPost, "/api/usercourse", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", tokenJWT)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Create User Course 2
				requestBody = strings.NewReader(`{"user_id": 10,"course_id": 2}`)
				request = httptest.NewRequest(http.MethodPost, "/api/usercourse", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", tokenJWT)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Create User Course 3
				requestBody = strings.NewReader(`{"user_id": 11,"course_id": 1}`)
				request = httptest.NewRequest(http.MethodPost, "/api/usercourse", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", tokenJWT)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// find all user courses
				request = httptest.NewRequest(http.MethodGet, "/api/usercourse", nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", tokenJWT)

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
})
