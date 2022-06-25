package integration

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-12/backend/config"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/test/setup"
)

var _ = Describe("Question API", func() {
	var (
		server         *gin.Engine
		token          string
		ok             bool
		createquestion model.CreateQuestionRequest
	)

	BeforeEach(func() {
		configuration := config.New("../../.env.test")

		_, err := setup.SuiteSetup(configuration)
		if err != nil {
			panic(err)
		}

		router := setup.ModuleSetup(configuration)
		server = router

		createquestion = model.CreateQuestionRequest{
			UserId:      1,
			Title:       "Algoritma Naive Bayes",
			ModuleId:    1,
			Tags:        "#NaiveBayes",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		}

		var user = model.UserRegisterResponse{
			Name:           "akuntest",
			Username:       "akuntest",
			Email:          "akuntest@gmail.com",
			Password:       "123456ll",
			Role:           2,
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
		userData, _ = json.Marshal(login)
		requestBody = strings.NewReader(string(userData))
		request = httptest.NewRequest(http.MethodPost, "/api/users/login", requestBody)
		request.Header.Add("Content-Type", "application/json")

		writer = httptest.NewRecorder()
		server.ServeHTTP(writer, request)

		responseLogin := writer.Result()

		body, _ := io.ReadAll(responseLogin.Body)
		var responseBodyLogin map[string]interface{}
		_ = json.Unmarshal(body, &responseBodyLogin)

		log.Println(responseBodyLogin["status"])
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

	Describe("Create Question", func() {
		When("Question is empty", func() {
			It("Should return data", func() {
				questionData, _ := json.Marshal(createquestion)
				requestBody := strings.NewReader(string(questionData))
				request := httptest.NewRequest(http.MethodPost, "/api/questions/create", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				Expect(int(responseBody["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody["status"]).To(Equal("question successfully created"))
				Expect(responseBody["data"].(map[string]interface{})["user_id"]).To(Equal(float64(1)))
				Expect(responseBody["data"].(map[string]interface{})["module_id"]).To(Equal(float64(1)))
				Expect(responseBody["data"].(map[string]interface{})["title"]).To(Equal("Algoritma Naive Bayes"))
				Expect(responseBody["data"].(map[string]interface{})["tags"]).To(Equal("#NaiveBayes"))
				Expect(responseBody["data"].(map[string]interface{})["description"]).To(Equal("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."))
			})
		})
	})
	Describe("Create Question and Get All", func() {
		When("Question is empty", func() {
			It("Should return data", func() {
				questionData, _ := json.Marshal(createquestion)
				requestBody := strings.NewReader(string(questionData))
				request := httptest.NewRequest(http.MethodPost, "/api/questions/create", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Get All Question
				request = httptest.NewRequest(http.MethodGet, "/api/questions/all", nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				Expect(int(responseBody["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody["status"]).To(Equal("OK"))
			})
		})
	})
	Describe("Create Question and Get by user id", func() {
		When("Question is empty", func() {
			It("Should return data", func() {
				questionData, _ := json.Marshal(createquestion)
				requestBody := strings.NewReader(string(questionData))
				request := httptest.NewRequest(http.MethodPost, "/api/questions/create", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Get All Question
				request = httptest.NewRequest(http.MethodGet, "/api/questions/by-user/1", nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				Expect(int(responseBody["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody["status"]).To(Equal("OK"))
			})
		})
	})
	Describe("Update Question", func() {
		When("Question is empty", func() {
			It("Should return error because user id is not match", func() {
				var updatequestion = model.UpdateQuestionRequest{
					UserId:      1,
					Title:       "Algoritma Naive Bayes",
					ModuleId:    2,
					Tags:        "#NaiveBayes",
					Description: "labore et dolore magna aliqua.",
				}

				questionData, _ := json.Marshal(createquestion)
				requestBody := strings.NewReader(string(questionData))
				request := httptest.NewRequest(http.MethodPost, "/api/questions/create", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Update Question
				questionData, _ = json.Marshal(updatequestion)
				requestBody = strings.NewReader(string(questionData))
				request = httptest.NewRequest(http.MethodPut, "/api/questions/update/1", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				Expect(int(responseBody["code"].(float64))).To(Equal(500))
			})
		})
	})
	Describe("Delete Question", func() {
		When("Question is empty", func() {
			It("Should return error because user id is not match", func() {
				questionData, _ := json.Marshal(createquestion)
				requestBody := strings.NewReader(string(questionData))
				request := httptest.NewRequest(http.MethodPost, "/api/questions/create", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Delete Question
				request = httptest.NewRequest(http.MethodDelete, "/api/questions/1", nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				Expect(int(responseBody["code"].(float64))).To(Equal(500))
			})
		})
	})
})
