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

var _ = Describe("Courses API", func() {

	var (
		server       *gin.Engine
		token        string
		idUser       int
		nameUser     string
		usernameUser string
		emailUser    string
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
		nameUser = responseBodyRegister["data"].(map[string]interface{})["name"].(string)
		usernameUser = responseBodyRegister["data"].(map[string]interface{})["username"].(string)
		emailUser = responseBodyRegister["data"].(map[string]interface{})["email"].(string)

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

	Describe("Find All Courses", func() {
		When("the data is exists", func() {
			It("should return all courses response", func() {
				// Create Course 1
				requestBody := strings.NewReader(`{"name": "Teknik Komputer Jaringan","class": "TKJ-3","tools": "Router, RJ-45","about": "Pada pelajaran kali ini akan lebih difokuskan pada pembuatan tower","description": "Siswa mampu membuat tower sendiri"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Create Course 2
				requestBody = strings.NewReader(`{"name": "Rekayasa Perangkat Lunak","class": "RPL-1","tools": "XAMPP","about": "Pada pelajaran kali ini akan lebih difokuskan pada pembuatan tower","description": "Siswa mampu membuat tower sendiri"}`)
				request = httptest.NewRequest(http.MethodPost, "/api/courses", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Find All Course
				request = httptest.NewRequest(http.MethodGet, "/api/courses", nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				courses := responseBody["data"].([]interface{})
				courseResponse1 := courses[0].(map[string]interface{})
				courseResponse2 := courses[1].(map[string]interface{})

				Expect(int(responseBody["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody["status"]).To(Equal("OK"))

				Expect(courseResponse1["name"]).To(Equal("Teknik Komputer Jaringan"))
				Expect(courseResponse1["class"]).To(Equal("TKJ-3"))
				Expect(courseResponse1["tools"]).To(Equal("Router, RJ-45"))

				Expect(courseResponse2["name"]).To(Equal("Rekayasa Perangkat Lunak"))
				Expect(courseResponse2["class"]).To(Equal("RPL-1"))
				Expect(courseResponse2["tools"]).To(Equal("XAMPP"))
			})
		})
	})

	Describe("Create Courses", func() {
		When("the fields are filled", func() {
			It("should return successful create courses response", func() {
				// Create Course
				requestBody := strings.NewReader(`{"name": "Teknik Komputer Jaringan","class": "TKJ-3","tools": "Router, RJ-45","about": "Pada pelajaran kali ini akan lebih difokuskan pada pembuatan tower","description": "Siswa mampu membuat tower sendiri"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				Expect(int(responseBody["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody["status"]).To(Equal("course successfully created"))
				Expect(responseBody["data"].(map[string]interface{})["name"]).To(Equal("Teknik Komputer Jaringan"))
				Expect(responseBody["data"].(map[string]interface{})["class"]).To(Equal("TKJ-3"))
				Expect(responseBody["data"].(map[string]interface{})["tools"]).To(Equal("Router, RJ-45"))
				Expect(responseBody["data"].(map[string]interface{})["about"]).To(Equal("Pada pelajaran kali ini akan lebih difokuskan pada pembuatan tower"))
				Expect(responseBody["data"].(map[string]interface{})["description"]).To(Equal("Siswa mampu membuat tower sendiri"))
				Expect(responseBody["data"].(map[string]interface{})["is_active"]).To(BeTrue())
			})
		})
	})

	Describe("Find Course By Code", func() {
		When("the data is exists", func() {
			It("should return one course response", func() {
				// Create Course
				requestBody := strings.NewReader(`{"name": "Teknik Komputer Jaringan","class": "TKJ-3","tools": "Router, RJ-45","about": "Pada pelajaran kali ini akan lebih difokuskan pada pembuatan tower","description": "Siswa mampu membuat tower sendiri"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				// Find Course By Code
				codeCourse := responseBody["data"].(map[string]interface{})["code_course"].(string)
				request = httptest.NewRequest(http.MethodGet, "/api/courses/"+codeCourse, nil)
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
				Expect(responseBody1["data"].(map[string]interface{})["name"]).To(Equal("Teknik Komputer Jaringan"))
				Expect(responseBody1["data"].(map[string]interface{})["code_course"]).To(Equal(codeCourse))
				Expect(responseBody1["data"].(map[string]interface{})["class"]).To(Equal("TKJ-3"))
				Expect(responseBody1["data"].(map[string]interface{})["tools"]).To(Equal("Router, RJ-45"))
				Expect(responseBody1["data"].(map[string]interface{})["about"]).To(Equal("Pada pelajaran kali ini akan lebih difokuskan pada pembuatan tower"))
				Expect(responseBody1["data"].(map[string]interface{})["description"]).To(Equal("Siswa mampu membuat tower sendiri"))
				Expect(responseBody1["data"].(map[string]interface{})["is_active"]).To(BeTrue())
			})
		})
	})

	Describe("Update Course By Code", func() {
		When("the data is exists", func() {
			It("should return successfully update course response", func() {
				// Create Course
				requestBody := strings.NewReader(`{"name": "Teknik Komputer Jaringan","class": "TKJ-3","tools": "Router, RJ-45","about": "Pada pelajaran kali ini akan lebih difokuskan pada pembuatan tower","description": "Siswa mampu membuat tower sendiri"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				// Update Course By Code
				requestBody = strings.NewReader(`{"name": "Matematika","class": "IPA-3","tools": "Pensil","about": "Belajar aljabar","description": "Siswa mampu membuat rumus sendiri"}`)
				codeCourse := responseBody["data"].(map[string]interface{})["code_course"].(string)
				request = httptest.NewRequest(http.MethodPatch, "/api/courses/"+codeCourse, requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response = writer.Result()

				body, _ = io.ReadAll(response.Body)
				var responseBody1 map[string]interface{}
				_ = json.Unmarshal(body, &responseBody1)

				Expect(int(responseBody1["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody1["status"]).To(Equal("course successfully updated"))
				Expect(responseBody1["data"].(map[string]interface{})["name"]).To(Equal("Matematika"))
				Expect(responseBody1["data"].(map[string]interface{})["code_course"]).To(Equal(codeCourse))
				Expect(responseBody1["data"].(map[string]interface{})["class"]).To(Equal("IPA-3"))
				Expect(responseBody1["data"].(map[string]interface{})["tools"]).To(Equal("Pensil"))
				Expect(responseBody1["data"].(map[string]interface{})["about"]).To(Equal("Belajar aljabar"))
				Expect(responseBody1["data"].(map[string]interface{})["description"]).To(Equal("Siswa mampu membuat rumus sendiri"))
				Expect(responseBody1["data"].(map[string]interface{})["is_active"]).To(BeTrue())
			})
		})
	})

	Describe("Delete Course By Code", func() {
		When("the data is exists", func() {
			It("should return successfully delete course response", func() {
				// Create Course
				requestBody := strings.NewReader(`{"name": "Teknik Komputer Jaringan","class": "TKJ-3","tools": "Router, RJ-45","about": "Pada pelajaran kali ini akan lebih difokuskan pada pembuatan tower","description": "Siswa mampu membuat tower sendiri"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				// Delete Course By Code
				codeCourse := responseBody["data"].(map[string]interface{})["code_course"].(string)
				request = httptest.NewRequest(http.MethodDelete, "/api/courses/"+codeCourse, nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response = writer.Result()

				body, _ = io.ReadAll(response.Body)
				var responseBody1 map[string]interface{}
				_ = json.Unmarshal(body, &responseBody1)

				Expect(int(responseBody1["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody1["status"]).To(Equal("course successfully deleted"))
			})
		})
	})

	Describe("Change Status Course By Code", func() {
		When("the data is exists", func() {
			It("should return successfully delete course response", func() {
				// Create Course
				requestBody := strings.NewReader(`{"name": "Teknik Komputer Jaringan","class": "TKJ-3","tools": "Router, RJ-45","about": "Pada pelajaran kali ini akan lebih difokuskan pada pembuatan tower","description": "Siswa mampu membuat tower sendiri"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				// Update Course By Code
				codeCourse := responseBody["data"].(map[string]interface{})["code_course"].(string)
				requestBody = strings.NewReader(`{"is_active": false}`)
				request = httptest.NewRequest(http.MethodPatch, "/api/courses/"+codeCourse+"/status", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response = writer.Result()

				body, _ = io.ReadAll(response.Body)
				var responseBody1 map[string]interface{}
				_ = json.Unmarshal(body, &responseBody1)

				Expect(int(responseBody1["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody1["status"]).To(Equal("course successfully updated"))
			})
		})
	})

	Describe("Find All User By Code Course", func() {
		When("the data is exists", func() {
			It("should return successfully show all user's submissions response", func() {
				// Create Course
				requestBody := strings.NewReader(`{"name": "Teknik Komputer Jaringan","class": "TKJ-3","tools": "Router, RJ-45","about": "Pada pelajaran kali ini akan lebih difokuskan pada pembuatan tower","description": "Siswa mampu membuat tower sendiri"}`)
				request := httptest.NewRequest(http.MethodPost, "/api/courses", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				// Create User Course
				idCourse := int(responseBody["data"].(map[string]interface{})["id"].(float64))
				bodyUserCourse := fmt.Sprintf(`{"user_id": %v,"course_id": %v}`, idUser, idCourse)
				requestBody = strings.NewReader(bodyUserCourse)
				request = httptest.NewRequest(http.MethodPost, "/api/usercourse", requestBody)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Find All User By Code Course
				codeCourse := responseBody["data"].(map[string]interface{})["code_course"].(string)
				request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/courses/%v/users", codeCourse), nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response = writer.Result()

				body, _ = io.ReadAll(response.Body)
				var responseBody1 map[string]interface{}
				_ = json.Unmarshal(body, &responseBody1)

				items := responseBody1["data"].([]interface{})
				itemResponse := items[0].(map[string]interface{})

				Expect(int(responseBody1["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBody1["status"]).To(Equal("OK"))
				Expect(int(itemResponse["id_user"].(float64))).To(Equal(idUser))
				Expect(itemResponse["user_name"]).To(Equal(nameUser))
				Expect(itemResponse["user_username"]).To(Equal(usernameUser))
				Expect(itemResponse["user_email"]).To(Equal(emailUser))
			})
		})
	})
})
