package integration

import (
	"encoding/json"
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

var _ = Describe("User API", func() {

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

	var user = []model.UserRegisterResponse{{
		Name:           "usertest",
		Username:       "user",
		Email:          "user@gmail.com",
		Password:       "usertest",
		Role:           2,
		Phone:          "8131313131313",
		Gender:         2,
		DisabilityType: 1,
		Birthdate:      "2002-04-04",
	},
		{
			Name:           "useradmin",
			Username:       "admin",
			Email:          "admin@gmail.com",
			Password:       "useradmin",
			Role:           1,
			Phone:          "8121212121212",
			Gender:         2,
			DisabilityType: 1,
			Birthdate:      "2002-04-01",
		},
	}

	login := []model.GetUserLogin{
		{
			Email:    "user@gmail.com",
			Password: "usertest",
		},
		{
			Email:    "admin@gmail.com",
			Password: "useradmin",
		},
	}

	Describe("Register User", func() {
		When("the data is correctly initialized", func() {
			It("should return true if the user is registered", func() {
				// Register User
				userData, _ := json.Marshal(user[0])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				users := responseBody["data"]
				userResponse := users.(map[string]interface{})

				Expect(int(responseBody["code"].(float64))).To(Equal(http.StatusCreated))
				Expect(responseBody["status"]).To(Equal("User Register Successfully"))

				Expect(userResponse["name"]).To(Equal("usertest"))
				Expect(userResponse["username"]).To(Equal("user"))
				Expect(userResponse["email"]).To(Equal("user@gmail.com"))
				Expect(userResponse["password"]).ToNot(BeNil())
				Expect(userResponse["role"]).To(Equal(float64(2)))
				Expect(userResponse["phone"]).To(Equal("8131313131313"))
				Expect(userResponse["gender"]).To(Equal(float64(2)))
				Expect(userResponse["type_of_disability"]).To(Equal(float64(1)))
				Expect(userResponse["birthdate"]).To(Equal("2002-04-04"))
			})
		})

		When("Username and Email are registered", func() {
			It("Should return error", func() {
				// Register User
				userData, _ := json.Marshal(user[0])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				response := writer.Result()

				body, _ := io.ReadAll(response.Body)
				var responseBody map[string]interface{}
				_ = json.Unmarshal(body, &responseBody)

				userDataError, _ := json.Marshal(user[0])
				requestBodyError := strings.NewReader(string(userDataError))
				requestError := httptest.NewRequest(http.MethodPost, "/api/users", requestBodyError)
				request.Header.Add("Content-Type", "application/json")

				writerError := httptest.NewRecorder()
				server.ServeHTTP(writerError, requestError)

				responseError := writerError.Result()

				bodyError, _ := io.ReadAll(responseError.Body)
				var mapResponse map[string]interface{}
				_ = json.Unmarshal(bodyError, &mapResponse)

				Expect(int(mapResponse["code"].(float64))).To(Equal(http.StatusUnauthorized))
				Expect(mapResponse["status"]).To(Equal("username has been registered"))
			})
		})
	})

	Describe("Login User", func() {
		When("Username and Password Correct", func() {
			It("Should return true", func() {
				// Register User
				userData, _ := json.Marshal(user[1])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Login User
				userLogin, _ := json.Marshal(login[1])
				requestBodyLogin := strings.NewReader(string(userLogin))
				requestLogin := httptest.NewRequest(http.MethodPost, "/api/users/login", requestBodyLogin)
				request.Header.Add("Content-Type", "application/json")

				writerLogin := httptest.NewRecorder()
				server.ServeHTTP(writerLogin, requestLogin)

				responseLogin := writerLogin.Result()

				bodyLogin, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(bodyLogin, &responseBodyLogin)

				token := responseBodyLogin["token"]
				users := responseBodyLogin["data"]
				userResponse := users.(map[string]interface{})

				Expect(int(responseBodyLogin["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyLogin["status"]).To(Equal("Login Successfull"))

				Expect(token).NotTo(BeNil())
				Expect(userResponse["name"]).To(Equal("useradmin"))
				Expect(userResponse["username"]).To(Equal("admin"))
				Expect(userResponse["email"]).To(Equal("admin@gmail.com"))
				Expect(userResponse["role"]).To(Equal(float64(1)))
				Expect(userResponse["gender"]).To(Equal(float64(2)))
				Expect(userResponse["type_of_disability"]).To(Equal(float64(1)))
			})
		})

		When("Username and Password Incorrect", func() {
			It("Should return error", func() {
				// Register User
				userData, _ := json.Marshal(user[1])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Login User
				userLogin, _ := json.Marshal(model.GetUserLogin{
					Email:    "admin@gmail.com.com",
					Password: "halo ges",
				})
				requestBodyLogin := strings.NewReader(string(userLogin))
				requestLogin := httptest.NewRequest(http.MethodPost, "/api/users/login", requestBodyLogin)
				request.Header.Add("Content-Type", "application/json")

				writerLogin := httptest.NewRecorder()
				server.ServeHTTP(writerLogin, requestLogin)

				responseLogin := writerLogin.Result()

				bodyLogin, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(bodyLogin, &responseBodyLogin)

				Expect(int(responseBodyLogin["code"].(float64))).To(Equal(http.StatusBadRequest))
				Expect(responseBodyLogin["status"]).To(Equal("Bad Request"))
				Expect(responseBodyLogin["data"]).To(Equal("Please Check Your Input"))
			})
		})
	})

	Describe("Check User Status", func() {
		When("User is logged in", func() {
			It("returns user is logged in", func() {
				// Register User
				userData, _ := json.Marshal(user[1])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				//Login User
				userLogin, _ := json.Marshal(login[1])
				requestBodyLogin := strings.NewReader(string(userLogin))
				requestLogin := httptest.NewRequest(http.MethodPost, "/api/users/login", requestBodyLogin)
				request.Header.Add("Content-Type", "application/json")

				writerLogin := httptest.NewRecorder()
				server.ServeHTTP(writerLogin, requestLogin)

				responseLogin := writerLogin.Result()

				bodyLogin, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(bodyLogin, &responseBodyLogin)

				token := responseBodyLogin["token"].(string)

				// Get User Status
				requestStatus := httptest.NewRequest(http.MethodGet, "/api/userstatus", nil)
				requestStatus.Header.Set("Content-Type", "application/json")
				requestStatus.Header.Set("Authorization", token)

				writerStatus := httptest.NewRecorder()
				server.ServeHTTP(writerStatus, requestStatus)

				responseStatus := writerStatus.Result()

				bodyStatus, _ := io.ReadAll(responseStatus.Body)
				var responseBodyStatus map[string]interface{}
				_ = json.Unmarshal(bodyStatus, &responseBodyStatus)

				userstatus := responseBodyStatus["data"]
				userResponseStatus := userstatus.(map[string]interface{})

				Expect(int(responseBodyStatus["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyStatus["status"]).To(Equal("User Already Logged In"))

				Expect(userResponseStatus["name"]).To(Equal("useradmin"))
				Expect(userResponseStatus["username"]).To(Equal("admin"))
				Expect(userResponseStatus["role"]).To(Equal(float64(1)))
				Expect(userResponseStatus["phone"]).To(Equal("8121212121212"))
				Expect(userResponseStatus["gender"]).To(Equal(float64(2)))
				Expect(userResponseStatus["type_of_disability"]).To(Equal(float64(1)))
			})
		})

		When("User is not logged in", func() {
			It("returns error", func() {
				// Register User
				userData, _ := json.Marshal(user[1])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Get User Status
				var token string
				requestStatus := httptest.NewRequest(http.MethodGet, "/api/userstatus", nil)
				requestStatus.Header.Set("Content-Type", "application/json")
				requestStatus.Header.Set("Authorization", token)

				writerStatus := httptest.NewRecorder()
				server.ServeHTTP(writerStatus, requestStatus)

				responseStatus := writerStatus.Result()

				bodyStatus, _ := io.ReadAll(responseStatus.Body)
				var responseBodyStatus map[string]interface{}
				_ = json.Unmarshal(bodyStatus, &responseBodyStatus)

				Expect(int(responseBodyStatus["code"].(float64))).To(Equal(http.StatusUnauthorized))
				Expect(responseBodyStatus["status"]).To(Equal("Unauthorized"))
			})
		})
	})

	Describe("Logout User", func() {
		When("User is logged in", func() {
			It("Return logout user successfull", func() {
				// Register User
				userData, _ := json.Marshal(user[0])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				//Login User
				userLogin, _ := json.Marshal(login[0])
				requestBodyLogin := strings.NewReader(string(userLogin))
				requestLogin := httptest.NewRequest(http.MethodPost, "/api/users/login", requestBodyLogin)
				request.Header.Add("Content-Type", "application/json")

				writerLogin := httptest.NewRecorder()
				server.ServeHTTP(writerLogin, requestLogin)

				responseLogin := writerLogin.Result()

				bodyLogin, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(bodyLogin, &responseBodyLogin)

				token := responseBodyLogin["token"].(string)

				// Logout User
				requestLogout := httptest.NewRequest(http.MethodPost, "/api/users/logout", nil)
				requestLogout.Header.Set("Content-Type", "application/json")
				requestLogout.Header.Set("Authorization", token)

				writerLogout := httptest.NewRecorder()
				server.ServeHTTP(writerLogout, requestLogout)

				responseLogout := writerLogout.Result()

				bodyLogout, _ := io.ReadAll(responseLogout.Body)
				var responseBodyLogout map[string]interface{}
				_ = json.Unmarshal(bodyLogout, &responseBodyLogout)

				Expect(int(responseBodyLogout["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyLogout["status"]).To(Equal("Logout Successfull"))
			})
		})

		When("User is logged in", func() {
			It("Return logout user successfull", func() {
				// Register User
				userData, _ := json.Marshal(user[0])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				// Logout User
				var token string
				requestLogout := httptest.NewRequest(http.MethodPost, "/api/users/logout", nil)
				requestLogout.Header.Set("Content-Type", "application/json")
				requestLogout.Header.Set("Authorization", token)

				writerLogout := httptest.NewRecorder()
				server.ServeHTTP(writerLogout, requestLogout)

				responseLogout := writerLogout.Result()

				bodyLogout, _ := io.ReadAll(responseLogout.Body)
				var responseBodyLogout map[string]interface{}
				_ = json.Unmarshal(bodyLogout, &responseBodyLogout)

				Expect(int(responseBodyLogout["code"].(float64))).To(Equal(http.StatusUnauthorized))
				Expect(responseBodyLogout["status"]).To(Equal("Unauthorized"))
			})
		})
	})

	Describe("List Users", func() {
		When("User is Admin", func() {
			It("Return all users", func() {
				// Register User
				userData, _ := json.Marshal(user[1])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				//Login User
				userLogin, _ := json.Marshal(login[1])
				requestBodyLogin := strings.NewReader(string(userLogin))
				requestLogin := httptest.NewRequest(http.MethodPost, "/api/users/login", requestBodyLogin)
				request.Header.Add("Content-Type", "application/json")

				writerLogin := httptest.NewRecorder()
				server.ServeHTTP(writerLogin, requestLogin)

				responseLogin := writerLogin.Result()

				bodyLogin, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(bodyLogin, &responseBodyLogin)

				token := responseBodyLogin["token"].(string)

				// List User
				requestList := httptest.NewRequest(http.MethodGet, "/api/users", nil)
				requestList.Header.Set("Content-Type", "application/json")
				requestList.Header.Set("Authorization", token)

				writerList := httptest.NewRecorder()
				server.ServeHTTP(writerList, requestList)

				responseList := writerList.Result()

				bodyList, _ := io.ReadAll(responseList.Body)
				var responseBodyList map[string]interface{}
				_ = json.Unmarshal(bodyList, &responseBodyList)

				userList := responseBodyList["data"].([]interface{})
				userResponseList := userList[0].(map[string]interface{})

				Expect(int(responseBodyList["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyList["status"]).To(Equal("Get All User Successfull"))

				Expect(userResponseList["name"]).To(Equal("useradmin"))
				Expect(userResponseList["username"]).To(Equal("admin"))
				Expect(userResponseList["role"]).To(Equal(float64(1)))
				Expect(userResponseList["phone"]).To(Equal("8121212121212"))
				Expect(userResponseList["gender"]).To(Equal(float64(2)))
				Expect(userResponseList["type_of_disability"]).To(Equal(float64(1)))
				Expect(userResponseList["birthdate"]).To(Equal("2002-04-01T00:00:00Z"))
			})
		})

		When("User is Not Admin", func() {
			It("Return Unauthorized", func() {
				// Register User
				userData, _ := json.Marshal(user[0])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				//Login User
				userLogin, _ := json.Marshal(login[0])
				requestBodyLogin := strings.NewReader(string(userLogin))
				requestLogin := httptest.NewRequest(http.MethodPost, "/api/users/login", requestBodyLogin)
				request.Header.Add("Content-Type", "application/json")

				writerLogin := httptest.NewRecorder()
				server.ServeHTTP(writerLogin, requestLogin)

				responseLogin := writerLogin.Result()

				bodyLogin, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(bodyLogin, &responseBodyLogin)

				token := responseBodyLogin["token"].(string)

				// List User
				requestList := httptest.NewRequest(http.MethodGet, "/api/users", nil)
				requestList.Header.Set("Content-Type", "application/json")
				requestList.Header.Set("Authorization", token)

				writerList := httptest.NewRecorder()
				server.ServeHTTP(writerList, requestList)

				responseList := writerList.Result()

				bodyList, _ := io.ReadAll(responseList.Body)
				var responseBodyList map[string]interface{}
				_ = json.Unmarshal(bodyList, &responseBodyList)

				Expect(int(responseBodyList["code"].(float64))).To(Equal(http.StatusUnauthorized))
				Expect(responseBodyList["status"]).To(Equal("You are not admin"))
			})
		})
	})
})
