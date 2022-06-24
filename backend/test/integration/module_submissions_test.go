package integration

import (
	"encoding/json"
	"fmt"
	"io"
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

var _ = Describe("Module Submissions API", func() {

	var (
		server     *gin.Engine
		token      string
		codeCourse string
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

		token = responseBodyLogin["token"].(string)

		// Create Course 1
		requestBody = strings.NewReader(`{"name": "Teknik Komputer Jaringan","class": "TKJ-3","tools": "Router, RJ-45","about": "Pada pelajaran kali ini akan lebih difokuskan pada pembuatan tower","description": "Siswa mampu membuat tower sendiri"}`)
		request = httptest.NewRequest(http.MethodPost, "/api/courses", requestBody)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Set("Authorization", token)

		writer = httptest.NewRecorder()
		server.ServeHTTP(writer, request)

		response := writer.Result()

		body, _ = io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		_ = json.Unmarshal(body, &responseBody)

		codeCourse = responseBody["data"].(map[string]interface{})["code_course"].(string)
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

	Describe("Find All Module Submission By Code Course", func() {
		When("the data is exists", func() {
			It("should return all module submission response", func() {
				// Create Module Submission
				requestBody := strings.NewReader(`{"name": "tugas Olahraga Bang","description": "renang","deadline": "2022-06-21T15:21:38+07:00"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses/"+codeCourse+"/submissions", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Find All Module Submission
				request = httptest.NewRequest(http.MethodGet, "/api/courses/"+codeCourse+"/submissions", nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				items := responseBody["data"].([]interface{})
				itemResponse := items[0].(map[string]interface{})

				Expect(int(responseBody["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody["status"]).To(Equal("OK"))
				Expect(itemResponse["name"]).To(Equal("tugas Olahraga Bang"))
				Expect(itemResponse["description"]).To(Equal("renang"))
				Expect(itemResponse["deadline"]).To(Equal("2022-06-21T15:21:38+07:00"))
			})
		})
	})

	Describe("Create Module Submission", func() {
		When("the fields are filled", func() {
			It("should return successful create module submission response", func() {
				// Create Module Submission
				requestBody := strings.NewReader(`{"name": "tugas Olahraga Bang","description": "renang","deadline": "2022-06-21T15:21:38+07:00"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses/"+codeCourse+"/submissions", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				Expect(int(responseBody["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody["status"]).To(Equal("module submission successfully created"))
				Expect(responseBody["data"].(map[string]interface{})["name"]).To(Equal("tugas Olahraga Bang"))
				Expect(responseBody["data"].(map[string]interface{})["description"]).To(Equal("renang"))
				Expect(responseBody["data"].(map[string]interface{})["deadline"]).To(Equal("2022-06-21T15:21:38+07:00"))
			})
		})
	})

	Describe("Find Module Submission By Id", func() {
		When("the data is exists", func() {
			It("should return one module submission response", func() {
				// Create Module Submission
				requestBody := strings.NewReader(`{"name": "tugas Olahraga Bang","description": "renang","deadline": "2022-06-21T15:21:38+07:00"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses/"+codeCourse+"/submissions", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				// Find Module Submission By Id
				id := int(responseBody["data"].(map[string]interface{})["id"].(float64))
				request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/courses/%v/submissions/%v", codeCourse, id), nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response = writer.Result()

				body, _ = io.ReadAll(response.Body)
				var responseBody1 map[string]interface{}
				_ = json.Unmarshal(body, &responseBody1)

				Expect(int(responseBody1["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody1["status"]).To(Equal("OK"))
				Expect(responseBody["data"].(map[string]interface{})["name"]).To(Equal("tugas Olahraga Bang"))
				Expect(responseBody["data"].(map[string]interface{})["description"]).To(Equal("renang"))
				Expect(responseBody["data"].(map[string]interface{})["deadline"]).To(Equal("2022-06-21T15:21:38+07:00"))
			})
		})
	})

	Describe("Update Module Submission By Id", func() {
		When("the data is exists", func() {
			It("should return successfully update module submission response", func() {
				// Create Module Submission
				requestBody := strings.NewReader(`{"name": "tugas Olahraga Bang","description": "renang","deadline": "2022-06-21T15:21:38+07:00"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses/"+codeCourse+"/submissions", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				// Update Module Submission By Id
				requestBody = strings.NewReader(`{"name": "Tugas Aljabar Bang","description": "harus jago MTK","deadline": "2022-07-21T15:21:38+07:00"}`)
				id := int(responseBody["data"].(map[string]interface{})["id"].(float64))
				request = httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/courses/%v/submissions/%v", codeCourse, id), requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response = writer.Result()

				body, _ = io.ReadAll(response.Body)
				var responseBody1 map[string]interface{}
				_ = json.Unmarshal(body, &responseBody1)

				Expect(int(responseBody1["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody1["status"]).To(Equal("module submission successfully updated"))
				Expect(responseBody1["data"].(map[string]interface{})["name"]).To(Equal("Tugas Aljabar Bang"))
				Expect(responseBody1["data"].(map[string]interface{})["description"]).To(Equal("harus jago MTK"))
				Expect(responseBody1["data"].(map[string]interface{})["deadline"]).To(Equal("2022-07-21T15:21:38+07:00"))
			})
		})
	})

	Describe("Delete Module Submission By Id", func() {
		When("the data is exists", func() {
			It("should return successfully delete module submission response", func() {
				// Create Module Submission
				requestBody := strings.NewReader(`{"name": "tugas Olahraga Bang","description": "renang","deadline": "2022-06-21T15:21:38+07:00"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses/"+codeCourse+"/submissions", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				// Delete Module Submission By Id
				id := int(responseBody["data"].(map[string]interface{})["id"].(float64))
				request = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/courses/%v/submissions/%v", codeCourse, id), nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response = writer.Result()

				body, _ = io.ReadAll(response.Body)
				var responseBody1 map[string]interface{}
				_ = json.Unmarshal(body, &responseBody1)

				Expect(int(responseBody1["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody1["status"]).To(Equal("module submission successfully deleted"))
			})
		})
	})

	Describe("Get Next Module Submission", func() {
		When("the data is exists", func() {
			It("should return successfully get next module submission response", func() {
				// Create Module Submission 1
				requestBody := strings.NewReader(`{"name": "tugas Olahraga Bang","description": "renang","deadline": "2022-06-21T15:21:38+07:00"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses/"+codeCourse+"/submissions", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				// Create Module Submission 2
				requestBody = strings.NewReader(`{"name": "tugas Olahraga Bang","description": "renang","deadline": "2022-06-21T15:21:38+07:00"}`)
				request = httptest.NewRequest(http.MethodPost, "/api/courses/"+codeCourse+"/submissions", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response = writer.Result()

				body, _ = io.ReadAll(response.Body)
				var responseBody1 map[string]interface{}
				_ = json.Unmarshal(body, &responseBody1)

				// Get Next Module Submission
				id := int(responseBody["data"].(map[string]interface{})["id"].(float64))
				request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/courses/%v/submissions/%v/next", codeCourse, id), nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response = writer.Result()

				body, _ = io.ReadAll(response.Body)
				var responseBody2 map[string]interface{}
				_ = json.Unmarshal(body, &responseBody2)

				Expect(int(responseBody2["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody2["status"]).To(Equal("OK"))
				Expect(int(responseBody2["data"].(map[string]interface{})["id"].(float64))).To(Equal(int(responseBody1["data"].(map[string]interface{})["id"].(float64))))
				Expect(responseBody2["data"].(map[string]interface{})["code_course"]).To(Equal(codeCourse))
			})
		})
	})

	Describe("Get Previous Module Submission", func() {
		When("the data is exists", func() {
			It("should return successfully get previous module submission response", func() {
				// Create Module Submission 1
				requestBody := strings.NewReader(`{"name": "tugas Olahraga Bang","description": "renang","deadline": "2022-06-21T15:21:38+07:00"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses/"+codeCourse+"/submissions", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				// Create Module Submission 2
				requestBody = strings.NewReader(`{"name": "tugas Olahraga Bang","description": "renang","deadline": "2022-06-21T15:21:38+07:00"}`)
				request = httptest.NewRequest(http.MethodPost, "/api/courses/"+codeCourse+"/submissions", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response = writer.Result()

				body, _ = io.ReadAll(response.Body)
				var responseBody1 map[string]interface{}
				_ = json.Unmarshal(body, &responseBody1)

				// Get Next Module Submission
				id := int(responseBody1["data"].(map[string]interface{})["id"].(float64))
				request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/courses/%v/submissions/%v/previous", codeCourse, id), nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response = writer.Result()

				body, _ = io.ReadAll(response.Body)
				var responseBody2 map[string]interface{}
				_ = json.Unmarshal(body, &responseBody2)

				Expect(int(responseBody2["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody2["status"]).To(Equal("OK"))
				Expect(int(responseBody2["data"].(map[string]interface{})["id"].(float64))).To(Equal(int(responseBody["data"].(map[string]interface{})["id"].(float64))))
				Expect(responseBody2["data"].(map[string]interface{})["code_course"]).To(Equal(codeCourse))
			})
		})
	})
})
