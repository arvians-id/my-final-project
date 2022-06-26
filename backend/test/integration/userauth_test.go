package integration

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	user := []model.UserRegisterResponse{{
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

	update := []model.GetUserDetailUpdate{
		{
			Name:           "usertest",
			Username:       "user",
			Role:           2,
			Phone:          "8131313131313",
			Gender:         2,
			DisabilityType: 1,
			Address:        "rumah orang",
			Birthdate:      "2002-04-04",
			Image:          "foto.png",
			Description:    "halo ges",
		},
		{
			Name:           "useradmin",
			Username:       "admin",
			Role:           1,
			Phone:          "8121212121212",
			Gender:         2,
			DisabilityType: 1,
			Address:        "ruang guru",
			Birthdate:      "2002-04-01",
			Image:          "fotoadmin.png",
			Description:    "ini akun admin",
		},
	}

	course := []model.CreateCourseRequest{
		{
			Name:        "Biologi",
			Class:       "12 IPA 1",
			Tools:       "Buku Paket, LKS, Buku Tulis",
			About:       "Ini adalah mata pelajaran biologi untuk siswa kelas 12 IPA 1",
			Description: "Siswa mampu memahami struktur jaringan pada tumbuhan dan diharapkan siswa dapat mengerjakan semua tugas dengan sebaik-baiknya",
		},
		{
			Name:        "Matematika",
			Class:       "12 IPA 2",
			Tools:       "Buku Paket, LKS, Buku Tulis",
			About:       "Ini adalah mata pelajaran matematika untuk siswa kelas 12 IPA 2",
			Description: "Siswa mampu memahami persamaan linear tiga variabel dan diharapkan siswa dapat mengerjakan semua tugas dengan sebaik-baiknya",
		},
	}

	Describe("Register and Login User", func() {
		When("The data is correct", func() {
			It("should return true if the user is registered and login", func() {
				// Register User
				userData, _ := json.Marshal(user[1])
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

				Expect(userResponse["name"]).To(Equal("useradmin"))
				Expect(userResponse["username"]).To(Equal("admin"))
				Expect(userResponse["email"]).To(Equal("admin@gmail.com"))
				Expect(userResponse["password"]).ToNot(BeNil())
				Expect(userResponse["role"]).To(Equal(float64(1)))
				Expect(userResponse["phone"]).To(Equal("8121212121212"))
				Expect(userResponse["gender"]).To(Equal(float64(2)))
				Expect(userResponse["type_of_disability"]).To(Equal(float64(1)))
				Expect(userResponse["birthdate"]).To(Equal("2002-04-01"))

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
				usersLogin := responseBodyLogin["data"]
				userResponseLogin := usersLogin.(map[string]interface{})

				Expect(int(responseBodyLogin["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyLogin["status"]).To(Equal("Login Successfull"))

				Expect(token).NotTo(BeNil())
				Expect(userResponseLogin["name"]).To(Equal("useradmin"))
				Expect(userResponseLogin["username"]).To(Equal("admin"))
				Expect(userResponseLogin["email"]).To(Equal("admin@gmail.com"))
				Expect(userResponseLogin["role"]).To(Equal(float64(1)))
				Expect(userResponseLogin["gender"]).To(Equal(float64(2)))
				Expect(userResponseLogin["type_of_disability"]).To(Equal(float64(1)))
			})
		})
	})

	Describe("Check User Status and Logout User", func() {
		When("User is logged in", func() {
			It("returns user is logged in and logged out", func() {
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
	})

	Describe("User can Get User By ID and Update the User Data", func() {
		When("User is authenticated", func() {
			It("Should return user details", func() {
				// Register User
				userData, _ := json.Marshal(user[0])
				requestBodyUser := strings.NewReader(string(userData))
				requestUser := httptest.NewRequest(http.MethodPost, "/api/users", requestBodyUser)
				requestUser.Header.Add("Content-Type", "application/json")

				writerUser := httptest.NewRecorder()
				server.ServeHTTP(writerUser, requestUser)

				//Login User
				adminLogin, _ := json.Marshal(login[0])
				requestBodyLogin := strings.NewReader(string(adminLogin))
				requestLogin := httptest.NewRequest(http.MethodPost, "/api/users/login", requestBodyLogin)
				requestLogin.Header.Add("Content-Type", "application/json")

				writerLogin := httptest.NewRecorder()
				server.ServeHTTP(writerLogin, requestLogin)

				responseLogin := writerLogin.Result()

				bodyLogin, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(bodyLogin, &responseBodyLogin)

				token := responseBodyLogin["token"].(string)

				// Update Role User
				tokenClaims := jwt.MapClaims{}
				jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
					return []byte("your secret api key"), nil
				},
				)
				idAdmin := tokenClaims["id"].(float64)
				idString := strconv.Itoa(int(idAdmin))
				requestDetailUser := httptest.NewRequest(http.MethodGet, "/api/users/"+idString, nil)
				requestDetailUser.Header.Set("Content-Type", "application/json")
				requestDetailUser.Header.Set("Authorization", token)

				writerDetailUser := httptest.NewRecorder()
				server.ServeHTTP(writerDetailUser, requestDetailUser)

				responseDetailUser := writerDetailUser.Result()

				bodyDetailUser, _ := io.ReadAll(responseDetailUser.Body)
				var responseBodyDetailUser map[string]interface{}
				_ = json.Unmarshal(bodyDetailUser, &responseBodyDetailUser)

				userDetailUser := responseBodyDetailUser["data"]
				userResponseDetailUser := userDetailUser.(map[string]interface{})

				Expect(int(responseBodyDetailUser["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyDetailUser["status"]).To(Equal("Get User By ID Successfull"))

				Expect(userResponseDetailUser["name"]).To(Equal("usertest"))
				Expect(userResponseDetailUser["username"]).To(Equal("user"))
				Expect(userResponseDetailUser["role"]).To(Equal(float64(2)))
				Expect(userResponseDetailUser["phone"]).To(Equal("8131313131313"))
				Expect(userResponseDetailUser["gender"]).To(Equal(float64(2)))
				Expect(userResponseDetailUser["type_of_disability"]).To(Equal(float64(1)))
				Expect(userResponseDetailUser["birthdate"]).To(Equal("2002-04-04T00:00:00Z"))

				userUpdate, _ := json.Marshal(update[0])
				requestBodyUpdate := strings.NewReader(string(userUpdate))
				requestUpdate := httptest.NewRequest(http.MethodPut, "/api/users/"+idString, requestBodyUpdate)
				requestUpdate.Header.Add("Content-Type", "application/json")
				requestUpdate.Header.Add("Authorization", token)

				writerUpdateUser := httptest.NewRecorder()
				server.ServeHTTP(writerUpdateUser, requestUpdate)

				responseUpdateUser := writerUpdateUser.Result()

				bodyUpdateUser, _ := io.ReadAll(responseUpdateUser.Body)
				var responseBodyUpdateUser map[string]interface{}
				_ = json.Unmarshal(bodyUpdateUser, &responseBodyUpdateUser)

				userUpdateUser := responseBodyUpdateUser["data"]
				userResponseUpdateUser := userUpdateUser.(map[string]interface{})

				Expect(int(responseBodyUpdateUser["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyUpdateUser["status"]).To(Equal("Update User Successfull"))

				Expect(userResponseUpdateUser["name"]).To(Equal("usertest"))
				Expect(userResponseUpdateUser["username"]).To(Equal("user"))
				Expect(userResponseUpdateUser["role"]).To(Equal(float64(2)))
				Expect(userResponseUpdateUser["phone"]).To(Equal("8131313131313"))
				Expect(userResponseUpdateUser["gender"]).To(Equal(float64(2)))
				Expect(userResponseUpdateUser["type_of_disability"]).To(Equal(float64(1)))
				Expect(userResponseUpdateUser["address"]).To(Equal("rumah orang"))
				Expect(userResponseUpdateUser["birthdate"]).To(Equal("2002-04-04"))
				Expect(userResponseUpdateUser["image"]).To(Equal("foto.png"))
				Expect(userResponseUpdateUser["description"]).To(Equal("halo ges"))
			})
		})
	})

	Describe("Admin can view all Users and Update Role User", func() {
		When("User is Admin", func() {
			It("Return all users", func() {
				// Register User
				userData, _ := json.Marshal(user[0])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				adminData, _ := json.Marshal(user[1])
				requestBodyAdmin := strings.NewReader(string(adminData))
				requestAdmin := httptest.NewRequest(http.MethodPost, "/api/users", requestBodyAdmin)
				requestAdmin.Header.Add("Content-Type", "application/json")

				writerAdmin := httptest.NewRecorder()
				server.ServeHTTP(writerAdmin, requestAdmin)

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
				adminResponseList := userList[1].(map[string]interface{})

				Expect(int(responseBodyList["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyList["status"]).To(Equal("Get All User Successfull"))

				Expect(userResponseList["name"]).To(Equal("usertest"))
				Expect(userResponseList["username"]).To(Equal("user"))
				Expect(userResponseList["role"]).To(Equal(float64(2)))
				Expect(userResponseList["phone"]).To(Equal("8131313131313"))
				Expect(userResponseList["gender"]).To(Equal(float64(2)))
				Expect(userResponseList["type_of_disability"]).To(Equal(float64(1)))
				Expect(userResponseList["birthdate"]).To(Equal("2002-04-04T00:00:00Z"))
				Expect(adminResponseList["name"]).To(Equal("useradmin"))
				Expect(adminResponseList["username"]).To(Equal("admin"))
				Expect(adminResponseList["role"]).To(Equal(float64(1)))
				Expect(adminResponseList["phone"]).To(Equal("8121212121212"))
				Expect(adminResponseList["gender"]).To(Equal(float64(2)))
				Expect(adminResponseList["type_of_disability"]).To(Equal(float64(1)))
				Expect(adminResponseList["birthdate"]).To(Equal("2002-04-01T00:00:00Z"))

				// Update Role User
				tokenClaims := jwt.MapClaims{}
				jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
					return []byte("your secret api key"), nil
				},
				)
				idAdmin := tokenClaims["id"].(float64)
				idAdmin = idAdmin - 1
				idString := strconv.Itoa(int(idAdmin))
				requestUpdate := httptest.NewRequest(http.MethodPut, "/api/users/roleupdate/"+idString, nil)
				requestUpdate.Header.Set("Content-Type", "application/json")
				requestUpdate.Header.Set("Authorization", token)

				writerUpdate := httptest.NewRecorder()
				server.ServeHTTP(writerUpdate, requestUpdate)

				responseUpdate := writerUpdate.Result()

				bodyUpdate, _ := io.ReadAll(responseUpdate.Body)
				var responseBodyUpdate map[string]interface{}
				_ = json.Unmarshal(bodyUpdate, &responseBodyUpdate)

				userUpdate := responseBodyUpdate["data"]
				userResponseUpdate := userUpdate.(map[string]interface{})

				Expect(int(responseBodyUpdate["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyUpdate["status"]).To(Equal("Update User Role Successfull"))

				Expect(userResponseUpdate["name"]).To(Equal("usertest"))
				Expect(userResponseUpdate["username"]).To(Equal("user"))
				Expect(userResponseUpdate["role"]).To(Equal(float64(1)))
				Expect(userResponseUpdate["phone"]).To(Equal("8131313131313"))
				Expect(userResponseUpdate["gender"]).To(Equal(float64(2)))
				Expect(userResponseUpdate["type_of_disability"]).To(Equal(float64(1)))
				Expect(userResponseUpdate["birthdate"]).To(Equal("2002-04-04T00:00:00Z"))
			})
		})
	})

	Describe("Admin can Delete User", func() {
		When("User is Admin", func() {
			It("Should return true", func() {
				// Register User
				userData, _ := json.Marshal(user[0])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				adminData, _ := json.Marshal(user[1])
				requestBodyAdmin := strings.NewReader(string(adminData))
				requestAdmin := httptest.NewRequest(http.MethodPost, "/api/users", requestBodyAdmin)
				requestAdmin.Header.Add("Content-Type", "application/json")

				writerAdmin := httptest.NewRecorder()
				server.ServeHTTP(writerAdmin, requestAdmin)

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

				// Delete User
				tokenClaims := jwt.MapClaims{}
				jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
					return []byte("your secret api key"), nil
				},
				)
				idAdmin := tokenClaims["id"].(float64)
				idAdmin = idAdmin - 1
				idString := strconv.Itoa(int(idAdmin))
				requestDelete := httptest.NewRequest(http.MethodDelete, "/api/users/"+idString, nil)
				requestDelete.Header.Set("Content-Type", "application/json")
				requestDelete.Header.Set("Authorization", token)

				writerDelete := httptest.NewRecorder()
				server.ServeHTTP(writerDelete, requestDelete)

				responseDelete := writerDelete.Result()

				bodyDelete, _ := io.ReadAll(responseDelete.Body)
				var responseBodyDelete map[string]interface{}
				_ = json.Unmarshal(bodyDelete, &responseBodyDelete)

				Expect(int(responseBodyDelete["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyDelete["status"]).To(Equal("Delete User Successfull"))
			})
		})
	})

	Describe("User can see their submissions", func() {
		When("User is logged in", func() {
			It("Should return all submissions", func() {
				// Register User
				userData, _ := json.Marshal(user[0])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				adminData, _ := json.Marshal(user[1])
				requestBodyAdmin := strings.NewReader(string(adminData))
				requestAdmin := httptest.NewRequest(http.MethodPost, "/api/users", requestBodyAdmin)
				requestAdmin.Header.Add("Content-Type", "application/json")

				writerAdmin := httptest.NewRecorder()
				server.ServeHTTP(writerAdmin, requestAdmin)

				//Login User
				userLogin, _ := json.Marshal(login[1])
				requestBodyLogin := strings.NewReader(string(userLogin))
				requestLogin := httptest.NewRequest(http.MethodPost, "/api/users/login", requestBodyLogin)
				requestLogin.Header.Add("Content-Type", "application/json")

				writerLogin := httptest.NewRecorder()
				server.ServeHTTP(writerLogin, requestLogin)

				responseLogin := writerLogin.Result()

				bodyLogin, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(bodyLogin, &responseBodyLogin)

				token := responseBodyLogin["token"].(string)

				// Delete User
				tokenClaims := jwt.MapClaims{}
				jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
					return []byte("your secret api key"), nil
				},
				)
				idAdmin := tokenClaims["id"].(float64)
				idString := strconv.Itoa(int(idAdmin))
				requestDelete := httptest.NewRequest(http.MethodDelete, "/api/users/"+idString, nil)
				requestDelete.Header.Set("Content-Type", "application/json")
				requestDelete.Header.Set("Authorization", token)

				writerDelete := httptest.NewRecorder()
				server.ServeHTTP(writerDelete, requestDelete)

				responseDelete := writerDelete.Result()

				bodyDelete, _ := io.ReadAll(responseDelete.Body)
				var responseBodyDelete map[string]interface{}
				_ = json.Unmarshal(bodyDelete, &responseBodyDelete)

				Expect(int(responseBodyDelete["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyDelete["status"]).To(Equal("Delete User Successfull"))
			})
		})
	})

	Describe("User can list user submissions", func() {
		When("User is logged in", func() {
			It("Should return user submission", func() {

				// Register User
				userData, _ := json.Marshal(user[0])
				requestBody := strings.NewReader(string(userData))
				request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
				request.Header.Add("Content-Type", "application/json")

				writer := httptest.NewRecorder()
				server.ServeHTTP(writer, request)

				adminData, _ := json.Marshal(user[1])
				requestBodyAdmin := strings.NewReader(string(adminData))
				requestAdmin := httptest.NewRequest(http.MethodPost, "/api/users", requestBodyAdmin)
				requestAdmin.Header.Add("Content-Type", "application/json")

				writerAdmin := httptest.NewRecorder()
				server.ServeHTTP(writerAdmin, requestAdmin)

				//Login User
				adminLogin, _ := json.Marshal(login[1])
				requestBodyLogin := strings.NewReader(string(adminLogin))
				requestLogin := httptest.NewRequest(http.MethodPost, "/api/users/login", requestBodyLogin)
				request.Header.Add("Content-Type", "application/json")

				writerLogin := httptest.NewRecorder()
				server.ServeHTTP(writerLogin, requestLogin)

				responseLogin := writerLogin.Result()

				bodyLogin, _ := io.ReadAll(responseLogin.Body)
				var responseBodyLogin map[string]interface{}
				_ = json.Unmarshal(bodyLogin, &responseBodyLogin)

				token := responseBodyLogin["token"].(string)

				// Create Course
				courseData1, _ := json.Marshal(course[0])
				requestBodyCourse1 := strings.NewReader(string(courseData1))
				requestCreateCourse1 := httptest.NewRequest(http.MethodPost, "/api/courses", requestBodyCourse1)
				requestCreateCourse1.Header.Add("Content-Type", "application/json")
				requestCreateCourse1.Header.Add("Authorization", token)

				writerCreateCourse1 := httptest.NewRecorder()
				server.ServeHTTP(writerCreateCourse1, requestCreateCourse1)

				courseData2, _ := json.Marshal(course[1])
				requestBodyCourse2 := strings.NewReader(string(courseData2))
				requestCreateCourse2 := httptest.NewRequest(http.MethodPost, "/api/courses", requestBodyCourse2)
				requestCreateCourse2.Header.Add("Content-Type", "application/json")
				requestCreateCourse2.Header.Add("Authorization", token)

				writerCreateCourse2 := httptest.NewRecorder()
				server.ServeHTTP(writerCreateCourse2, requestCreateCourse2)

				tokenClaims := jwt.MapClaims{}
				jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
					return []byte("your secret api key"), nil
				},
				)
				idUser := tokenClaims["id"].(float64)
				idUser = idUser - 1

				// List Course
				requestListCourse := httptest.NewRequest(http.MethodGet, "/api/courses", nil)
				requestListCourse.Header.Set("Content-Type", "application/json")
				requestListCourse.Header.Set("Authorization", token)

				writerList := httptest.NewRecorder()
				server.ServeHTTP(writerList, requestListCourse)

				responseList := writerList.Result()

				bodyList, _ := io.ReadAll(responseList.Body)
				var responseBodyList map[string]interface{}
				_ = json.Unmarshal(bodyList, &responseBodyList)

				courseList := responseBodyList["data"].([]interface{})
				courseResponseList1 := courseList[0].(map[string]interface{})
				courseResponseList2 := courseList[1].(map[string]interface{})

				idCourse1 := int(courseResponseList1["id"].(float64))
				codeCourse1 := courseResponseList1["code_course"].(string)
				idCourse2 := int(courseResponseList2["id"].(float64))
				codeCourse2 := courseResponseList2["code_course"].(string)

				userCourse := []model.CreateUserCourseRequest{
					{
						UserId:   int(idUser),
						CourseId: idCourse1,
					},
					{
						UserId:   int(idUser),
						CourseId: idCourse2,
					},
				}

				// Create User Course
				userCourse1, _ := json.Marshal(userCourse[0])
				requestBodyUserCourse1 := strings.NewReader(string(userCourse1))
				requestCreateUserCourse1 := httptest.NewRequest(http.MethodPost, "/api/usercourse", requestBodyUserCourse1)
				requestCreateUserCourse1.Header.Add("Content-Type", "application/json")
				requestCreateUserCourse1.Header.Add("Authorization", token)

				writerCreateUserCourse1 := httptest.NewRecorder()
				server.ServeHTTP(writerCreateUserCourse1, requestCreateUserCourse1)

				userCourse2, _ := json.Marshal(userCourse[0])
				requestBodyUserCourse2 := strings.NewReader(string(userCourse2))
				requestCreateUserCourse2 := httptest.NewRequest(http.MethodPost, "/api/usercourse", requestBodyUserCourse2)
				requestCreateUserCourse2.Header.Add("Content-Type", "application/json")
				requestCreateUserCourse2.Header.Add("Authorization", token)

				writerCreateUserCourse2 := httptest.NewRecorder()
				server.ServeHTTP(writerCreateUserCourse2, requestCreateUserCourse2)

				// Create User Submissions
				submissions := []model.CreateModuleSubmissionsRequest{
					{
						CourseId:    idCourse1,
						Name:        "Tugas Biologi",
						Description: "ini form untuk submissions biologi",
						Deadline:    time.Now().Add(2 * time.Hour),
					},
					{
						CourseId:    idCourse2,
						Name:        "Tugas Matematika",
						Description: "ini form untuk submissions matematika",
						Deadline:    time.Now().Add(2 * time.Hour),
					},
				}

				userSubmissions1, _ := json.Marshal(submissions[0])
				requestBodyUserSubmissions1 := strings.NewReader(string(userSubmissions1))
				requestCreateUserSubmissions1 := httptest.NewRequest(http.MethodPost, "/api/courses/"+codeCourse1+"/submissions", requestBodyUserSubmissions1)
				requestCreateUserSubmissions1.Header.Add("Content-Type", "application/json")
				requestCreateUserSubmissions1.Header.Add("Authorization", token)

				writerCreateUserSubmissions1 := httptest.NewRecorder()
				server.ServeHTTP(writerCreateUserSubmissions1, requestCreateUserSubmissions1)

				userSubmissions2, _ := json.Marshal(submissions[1])
				requestBodyUserSubmissions2 := strings.NewReader(string(userSubmissions2))
				requestCreateUserSubmissions2 := httptest.NewRequest(http.MethodPost, "/api/courses/"+codeCourse2+"/submissions", requestBodyUserSubmissions2)
				requestCreateUserSubmissions2.Header.Add("Content-Type", "application/json")
				requestCreateUserSubmissions2.Header.Add("Authorization", token)

				writerCreateUserSubmissions2 := httptest.NewRecorder()
				server.ServeHTTP(writerCreateUserSubmissions2, requestCreateUserSubmissions2)

				// Logout User
				requestLogout := httptest.NewRequest(http.MethodPost, "/api/users/logout", nil)
				requestLogout.Header.Set("Content-Type", "application/json")
				requestLogout.Header.Set("Authorization", token)

				writerLogout := httptest.NewRecorder()
				server.ServeHTTP(writerLogout, requestLogout)

				// Login User
				userLogin, _ := json.Marshal(login[0])
				requestBodyLoginUser := strings.NewReader(string(userLogin))
				requestLoginUser := httptest.NewRequest(http.MethodPost, "/api/users/login", requestBodyLoginUser)
				requestLoginUser.Header.Add("Content-Type", "application/json")

				writerLoginUser := httptest.NewRecorder()
				server.ServeHTTP(writerLoginUser, requestLoginUser)

				responseLoginUser := writerLoginUser.Result()

				bodyLoginUser, _ := io.ReadAll(responseLoginUser.Body)
				var responseBodyLoginUser map[string]interface{}
				_ = json.Unmarshal(bodyLoginUser, &responseBodyLoginUser)

				tokenUser := responseBodyLoginUser["token"].(string)

				// List User Submissions
				requestListSubmissions := httptest.NewRequest(http.MethodGet, "/api/users/submissions", nil)
				requestListSubmissions.Header.Set("Content-Type", "application/json")
				requestListSubmissions.Header.Set("Authorization", tokenUser)

				writerListSubmissions := httptest.NewRecorder()
				server.ServeHTTP(writerListSubmissions, requestListSubmissions)

				responseListSubmissions := writerListSubmissions.Result()

				bodyListSubmissions, _ := io.ReadAll(responseListSubmissions.Body)
				var responseBodyListSubmissions map[string]interface{}
				_ = json.Unmarshal(bodyListSubmissions, &responseBodyListSubmissions)

				userListSubmissions := responseBodyListSubmissions["data"].([]interface{})
				userResponseListSubmissions1 := userListSubmissions[0].(map[string]interface{})

				Expect(int(responseBodyListSubmissions["code"].(float64))).To(Equal(http.StatusOK))
				Expect(responseBodyListSubmissions["status"]).To(Equal("OK"))

				Expect(userResponseListSubmissions1["name_course"]).To(Equal("Biologi"))
				Expect(userResponseListSubmissions1["name_module_submission"]).To(Equal("Tugas Biologi"))
				Expect(userResponseListSubmissions1["grade"]).To(BeNil())
				Expect(userResponseListSubmissions1["file"]).To(BeNil())
			})
		})
	})
})
