package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-12/backend/config"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/test/setup"
)

var _ = Describe("User Submissions API", func() {

	var (
		server             *gin.Engine
		token              string
		codeCourse         string
		idUser             int
		idModuleSubmission int
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

		responseRegister := writer.Result()

		bodyRegister, _ := io.ReadAll(responseRegister.Body)
		var responseBodyRegister map[string]interface{}
		_ = json.Unmarshal(bodyRegister, &responseBodyRegister)

		idUser = int(responseBodyRegister["data"].(map[string]interface{})["id"].(float64))

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

		// Create Course
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
		idCourse := int(responseBody["data"].(map[string]interface{})["id"].(float64))

		// Create User Course
		rBody := fmt.Sprintf(`{"user_id":%v,"course_id":%v}`, idUser, idCourse)
		requestBody = strings.NewReader(rBody)
		request = httptest.NewRequest(http.MethodPost, "/api/usercourse", requestBody)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Set("Authorization", token)

		writer = httptest.NewRecorder()
		server.ServeHTTP(writer, request)

		// Create Module Submission
		requestBody = strings.NewReader(`{"name": "tugas Olahraga Bang","description": "renang","deadline": "2022-06-21T15:21:38+07:00"}`)
		request = httptest.NewRequest(http.MethodPost, "/api/courses/"+codeCourse+"/submissions", requestBody)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Set("Authorization", token)

		writer = httptest.NewRecorder()
		server.ServeHTTP(writer, request)

		response = writer.Result()

		body, _ = io.ReadAll(response.Body)
		var responseBodyModuleSubmission map[string]interface{}
		_ = json.Unmarshal(body, &responseBodyModuleSubmission)

		idModuleSubmission = int(responseBodyModuleSubmission["data"].(map[string]interface{})["id"].(float64))
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

	Describe("Submit File Submission", func() {
		When("the user submit file submission", func() {
			It("should return successful submit file submission response", func() {
				// Submit File Submission
				path := "./assets/example1.csv"
				body := new(bytes.Buffer)
				writer := multipart.NewWriter(body)
				part, _ := writer.CreateFormFile("file", path)
				sample, _ := os.Open(path)

				_, _ = io.Copy(part, sample)
				writer.Close()

				request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/courses/%v/submissions/%v/user-submit", codeCourse, idModuleSubmission), body)
				request.Header.Add("Content-Type", writer.FormDataContentType())
				request.Header.Set("Authorization", token)

				rec := httptest.NewRecorder()
				server.ServeHTTP(rec, request)

				response := rec.Result()

				resp, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(resp, &responseBody)

				log.Println(responseBody["status"])
				Expect(int(responseBody["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody["status"]).To(Equal("user submission successfully created"))
				Expect(int(responseBody["data"].(map[string]interface{})["user_id"].(float64))).To(Equal(idUser))
				Expect(int(responseBody["data"].(map[string]interface{})["module_submission_id"].(float64))).To(Equal(idModuleSubmission))
				Expect(responseBody["data"].(map[string]interface{})["file"].(string)).NotTo(BeNil())

				path, err := utils.GetPath("/assets/", responseBody["data"].(map[string]interface{})["file"].(string))
				if err != nil {
					panic(err)
				}

				err = os.Remove(path)
				if err != nil {
					panic(err)
				}
			})
		})
	})

	Describe("Find User File Submission By Id", func() {
		When("the file is exists", func() {
			It("should return successful get file submission response", func() {
				// Submit File Submission
				path := "./assets/example1.csv"
				body := new(bytes.Buffer)
				writer := multipart.NewWriter(body)
				part, _ := writer.CreateFormFile("file", path)
				sample, _ := os.Open(path)

				_, _ = io.Copy(part, sample)
				writer.Close()

				request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/courses/%v/submissions/%v/user-submit", codeCourse, idModuleSubmission), body)
				request.Header.Add("Content-Type", writer.FormDataContentType())
				request.Header.Set("Authorization", token)

				rec := httptest.NewRecorder()
				server.ServeHTTP(rec, request)

				response := rec.Result()

				resp, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(resp, &responseBody)

				// Find User File Submission By Id
				idUserSubmission := int(responseBody["data"].(map[string]interface{})["id"].(float64))
				request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/courses/%v/submissions/%v/user-submit/%v", codeCourse, idModuleSubmission, idUserSubmission), nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				rec = httptest.NewRecorder()
				server.ServeHTTP(rec, request)

				response = rec.Result()
				log.Println(response.Status)

				resp, _ = io.ReadAll(response.Body)
				var responseBody1 map[string]interface{}
				_ = json.Unmarshal(resp, &responseBody1)

				log.Println(responseBody1["status"])
				Expect(int(responseBody1["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody1["status"]).To(Equal("OK"))
				Expect(int(responseBody1["data"].(map[string]interface{})["id"].(float64))).To(Equal(idUserSubmission))
				Expect(int(responseBody1["data"].(map[string]interface{})["user_id"].(float64))).To(Equal(idUser))
				Expect(int(responseBody1["data"].(map[string]interface{})["module_submission_id"].(float64))).To(Equal(idModuleSubmission))
				Expect(responseBody1["data"].(map[string]interface{})["file"].(string)).To(Equal(responseBody["data"].(map[string]interface{})["file"].(string)))

				path, err := utils.GetPath("/assets/", responseBody["data"].(map[string]interface{})["file"].(string))
				if err != nil {
					panic(err)
				}

				err = os.Remove(path)
				if err != nil {
					panic(err)
				}
			})
		})
	})

	Describe("Update Grade User's Submission File", func() {
		When("the file is exists", func() {
			It("should return successful update grade user's file submission response", func() {
				// Submit File Submission
				path := "./assets/example1.csv"
				body := new(bytes.Buffer)
				writer := multipart.NewWriter(body)
				part, _ := writer.CreateFormFile("file", path)
				sample, _ := os.Open(path)

				_, _ = io.Copy(part, sample)
				writer.Close()

				request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/courses/%v/submissions/%v/user-submit", codeCourse, idModuleSubmission), body)
				request.Header.Add("Content-Type", writer.FormDataContentType())
				request.Header.Set("Authorization", token)

				rec := httptest.NewRecorder()
				server.ServeHTTP(rec, request)

				response := rec.Result()

				resp, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(resp, &responseBody)

				// Update Grade User's Submission File
				idUserSubmission := int(responseBody["data"].(map[string]interface{})["id"].(float64))
				requestBody := strings.NewReader(`{"grade": 60}`)
				request = httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/courses/%v/submissions/%v/user-submit/%v", codeCourse, idModuleSubmission, idUserSubmission), requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				rec = httptest.NewRecorder()
				server.ServeHTTP(rec, request)

				// Find User File Submission By Id
				request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/courses/%v/submissions/%v/user-submit/%v", codeCourse, idModuleSubmission, idUserSubmission), nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				rec = httptest.NewRecorder()
				server.ServeHTTP(rec, request)

				response = rec.Result()

				resp, _ = io.ReadAll(response.Body)
				var responseBody1 map[string]interface{}
				_ = json.Unmarshal(resp, &responseBody1)

				Expect(int(responseBody1["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody1["status"]).To(Equal("OK"))
				Expect(int(responseBody1["data"].(map[string]interface{})["id"].(float64))).To(Equal(idUserSubmission))
				Expect(int(responseBody1["data"].(map[string]interface{})["user_id"].(float64))).To(Equal(idUser))
				Expect(int(responseBody1["data"].(map[string]interface{})["module_submission_id"].(float64))).To(Equal(idModuleSubmission))
				Expect(responseBody1["data"].(map[string]interface{})["file"].(string)).To(Equal(responseBody["data"].(map[string]interface{})["file"].(string)))
				Expect(int(responseBody1["data"].(map[string]interface{})["grade"].(float64))).To(Equal(60))

				path, err := utils.GetPath("/assets/", responseBody["data"].(map[string]interface{})["file"].(string))
				if err != nil {
					panic(err)
				}

				err = os.Remove(path)
				if err != nil {
					panic(err)
				}
			})
		})
	})

	Describe("Download File Submission", func() {
		When("the file is exists", func() {
			It("should return successful download the file submission response", func() {
				// Submit File Submission
				path := "./assets/example1.csv"
				body := new(bytes.Buffer)
				writer := multipart.NewWriter(body)
				part, _ := writer.CreateFormFile("file", path)
				sample, _ := os.Open(path)

				_, _ = io.Copy(part, sample)
				writer.Close()

				request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/courses/%v/submissions/%v/user-submit", codeCourse, idModuleSubmission), body)
				request.Header.Add("Content-Type", writer.FormDataContentType())
				request.Header.Set("Authorization", token)

				rec := httptest.NewRecorder()
				server.ServeHTTP(rec, request)

				response := rec.Result()

				resp, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(resp, &responseBody)
				fileName := responseBody["data"].(map[string]interface{})["file"].(string)

				// Download File Submission
				log.Println(responseBody["status"])
				log.Println(responseBody["data"])
				idUserSubmission := int(responseBody["data"].(map[string]interface{})["id"].(float64))
				request = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/courses/%v/submissions/%v/user-submit/%v/download", codeCourse, idModuleSubmission, idUserSubmission), nil)
				contentDisposition := fmt.Sprintf("attachment; filename=%s", fileName)
				request.Header.Add("Content-Disposition", contentDisposition)
				request.Header.Set("Authorization", token)

				rec = httptest.NewRecorder()
				server.ServeHTTP(rec, request)

				response = rec.Result()

				Expect(response.StatusCode).To(Equal(http.StatusOK))

				path, err := utils.GetPath("/assets/", fileName)
				if err != nil {
					panic(err)
				}

				err = os.Remove(path)
				if err != nil {
					panic(err)
				}
			})
		})
	})
})
