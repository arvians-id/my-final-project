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

var _ = Describe("User Course API", func() {

	var (
		server *gin.Engine
		token  string
		ok     bool
		user   []model.UserRegisterResponse
		login  []model.GetUserLogin
	)

	BeforeEach(func() {
		configuration := config.New("../../.env.test")

		_, err := setup.SuiteSetup(configuration)
		if err != nil {
			panic(err)
		}

		router := setup.ModuleSetup(configuration)
		server = router

		user = []model.UserRegisterResponse{
			{
				Name:           "guru",
				Username:       "guru",
				Email:          "guru@gmail.com",
				Password:       "123456ll",
				Role:           1,
				Phone:          "085156789011",
				Gender:         1,
				DisabilityType: 1,
				Birthdate:      "2002-04-01",
			},
			{
				Name:           "murid",
				Username:       "murid",
				Email:          "murid@gmail.com",
				Password:       "123456ll",
				Role:           1,
				Phone:          "085156789011",
				Gender:         1,
				DisabilityType: 1,
				Birthdate:      "2002-04-01",
			},
		}

		login = []model.GetUserLogin{
			{
				Email:    "guru@gmail.com",
				Password: "123456ll",
			},
			{
				Email:    "murid@gmail.com",
				Password: "123456ll",
			},
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

	Describe("Create user guru and Login", func() {
		When("Data is empty", func() {
			It("Should return data", func() {
				// Register User
				userData, _ := json.Marshal(user[0])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				//Login User
				userData, _ = json.Marshal(login[0])
				requestBody = strings.NewReader(string(userData))
				request = httptest.NewRequest(http.MethodPost, "/api/users/login", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				responseLogin := writer.Result()

				body, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(body, &responseBodyLogin)

				token, ok = responseBodyLogin["token"].(string)
				if !ok {
					panic("Can't get token")
				} else {
					log.Println("Token: ", token)
				}

				Expect(int(responseBodyLogin["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyLogin["status"]).To(Equal("Login Successfull"))
				Expect(responseBodyLogin["token"]).To(Equal(token))
			})
		})
	})
	Describe("Create user murid and Login", func() {
		When("Data is empty", func() {
			It("Should return data", func() {
				// Register User
				userData, _ := json.Marshal(user[1])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				//Login User
				userData, _ = json.Marshal(login[1])
				requestBody = strings.NewReader(string(userData))
				request = httptest.NewRequest(http.MethodPost, "/api/users/login", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				responseLogin := writer.Result()

				body, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(body, &responseBodyLogin)

				token, ok = responseBodyLogin["token"].(string)
				if !ok {
					panic("Can't get token")
				} else {
					log.Println("Token: ", token)
				}

				Expect(int(responseBodyLogin["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyLogin["status"]).To(Equal("Login Successfull"))
				Expect(responseBodyLogin["token"]).To(Equal(token))
			})
		})
	})
	Describe("Login and Logout", func() {
		When("Data is empty", func() {
			It("Should return data", func() {
				// Register User
				userData, _ := json.Marshal(user[0])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				//Login User
				userData, _ = json.Marshal(login[0])
				requestBody = strings.NewReader(string(userData))
				request = httptest.NewRequest(http.MethodPost, "/api/users/login", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				responseLogin := writer.Result()

				body, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(body, &responseBodyLogin)

				token, ok = responseBodyLogin["token"].(string)
				if !ok {
					panic("Can't get token")
				} else {
					log.Println("Token: ", token)
				}

				Expect(int(responseBodyLogin["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyLogin["status"]).To(Equal("Login Successfull"))
				Expect(responseBodyLogin["token"]).To(Equal(token))

				// logout User
				request = httptest.NewRequest(http.MethodPost, "/api/users/logout", nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				responseLogout := writer.Result()

				body, _ = io.ReadAll(responseLogout.Body)
				var responseBodyLogout map[string]interface{}
				_ = json.Unmarshal(body, &responseBodyLogout)

				Expect(int(responseBodyLogout["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyLogout["status"]).To(Equal("Logout Successfull"))
			})
		})
	})
	Describe("Get user status murid", func() {
		When("Data is empty", func() {
			It("Should return data", func() {
				// Register User
				userData, _ := json.Marshal(user[1])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				//Login User
				userData, _ = json.Marshal(login[1])
				requestBody = strings.NewReader(string(userData))
				request = httptest.NewRequest(http.MethodPost, "/api/users/login", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				responseLogin := writer.Result()

				body, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(body, &responseBodyLogin)

				token, ok = responseBodyLogin["token"].(string)
				if !ok {
					panic("Can't get token")
				} else {
					log.Println("Token: ", token)
				}

				Expect(int(responseBodyLogin["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyLogin["status"]).To(Equal("Login Successfull"))
				Expect(responseBodyLogin["token"]).To(Equal(token))

				// Get user status murid
				request = httptest.NewRequest(http.MethodGet, "/api/userstatus", nil)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Set("Authorization", token)

				writer = httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				responseGetUserStatus := writer.Result()

				body, _ = io.ReadAll(responseGetUserStatus.Body)
				var responseBodyGetUserStatus map[string]interface{}
				_ = json.Unmarshal(body, &responseBodyGetUserStatus)

				Expect(int(responseBodyGetUserStatus["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyGetUserStatus["status"]).To(Equal("User Already Logged In"))
			})
		})
	})
})
