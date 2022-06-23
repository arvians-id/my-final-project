package unit_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	"github.com/rg-km/final-project-engineering-12/backend/config"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/test/setup"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("Test", func() {

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
})
