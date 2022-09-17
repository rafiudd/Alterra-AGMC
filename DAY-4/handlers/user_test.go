package handlers

import (
	"bytes"
	"DAY_4/config"
	"DAY_4/middleware"
	"DAY_4/mock"
	"DAY_4/models"
	"DAY_4/repositories"
	"DAY_4/services"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	var (
		repository  = repositories.NewRepositories(config.GetQuery())
		service     = services.NewServices(repository)
		h           = NewHandlers(service)
		echoMock    = mock.EchoMock{E: echo.New()}
		payloadUser = &models.User{
			Username: "test4",
			Email:    "test4@mail.com",
			Password: "123",
		}
		payloadUser2 = &models.User{
			Username: "test4",
			Email:    "test4mail.com",
			Password: "123",
		}
	)

	config.NewTrunc().DeleteDataUsers()
	t.Run("success", func(t *testing.T) {
		payload, err := json.Marshal(payloadUser)
		if err != nil {
			t.Fatal(err)
		}

		c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
		c.SetPath("/v1/users/")

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.CreateUser(c)) {
			asserts.Equal(201, rec.Code)
		}
	})

	t.Run("falied", func(t *testing.T) {
		payload, err := json.Marshal(payloadUser2)
		if err != nil {
			t.Fatal(err)
		}

		c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
		c.SetPath("/v1/users/")

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.CreateUser(c)) {
			asserts.Equal(400, rec.Code)
		}
	})
}

func TestLogin(t *testing.T) {
	t.Parallel()

	var (
		repository = repositories.NewRepositories(config.GetQuery())
		service    = services.NewServices(repository)
		h          = NewHandlers(service)
		echoMock   = mock.EchoMock{E: echo.New()}
		userJSON   = `{
			"password": "1234",
			"username": "test1"
		}`
	)

	config.NewTrunc().DeleteDataUsers()
	config.NewSeed().UsersSeeder()
	t.Run("success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/v1/login", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.UserLogin(c)) {
			asserts.Equal(200, rec.Code)
		}
	})

	t.Run("failed", func(t *testing.T) {
		c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
		c.SetPath("/v1/login")

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.UserLogin(c)) {
			asserts.Equal(400, rec.Code)
		}
	})
}

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	var (
		repository  = repositories.NewRepositories(config.GetQuery())
		service     = services.NewServices(repository)
		h           = NewHandlers(service)
		echoMock    = mock.EchoMock{E: echo.New()}
		userId      = int(1)
		payloadUser = &models.User{
			Username: "test",
			Email:    "test@mail.com",
		}
	)

	config.NewTrunc().DeleteDataUsers()
	config.NewSeed().UsersSeeder()
	t.Run("success", func(t *testing.T) {
		payload, err := json.Marshal(payloadUser)
		if err != nil {
			t.Fatal(err)
		}
		c, rec := echoMock.RequestMock(http.MethodGet, "/", bytes.NewBuffer(payload))
		c.SetPath("/jwt/v1/users:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.UpdateUser(c)) {
			asserts.Equal(200, rec.Code)
		}
	})

	t.Run("failed", func(t *testing.T) {
		payload, err := json.Marshal(payloadUser)
		if err != nil {
			t.Fatal(err)
		}
		c, rec := echoMock.RequestMock(http.MethodGet, "/", bytes.NewBuffer(payload))
		c.SetPath("/jwt/v1/users:id")
		c.SetParamNames("id")
		c.SetParamValues("200")

		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.UpdateUser(c)) {
			asserts.Equal(404, rec.Code)
		}
	})
}

func TestGetUser(t *testing.T) {
	t.Parallel()

	var (
		repository = repositories.NewRepositories(config.GetQuery())
		service    = services.NewServices(repository)
		h          = NewHandlers(service)
		echoMock   = mock.EchoMock{E: echo.New()}
		userId     = int(1)
	)

	config.NewTrunc().DeleteDataUsers()
	config.NewSeed().UsersSeeder()
	t.Run("success", func(t *testing.T) {
		c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
		c.SetPath("/jwt/v1/users:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.GetUserById(c)) {
			asserts.Equal(200, rec.Code)
		}
	})

	t.Run("failed", func(t *testing.T) {
		c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
		c.SetPath("/jwt/v1/users:id")
		c.SetParamNames("id")
		c.SetParamValues("200")

		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.GetUserById(c)) {
			asserts.Equal(404, rec.Code)
		}
	})
}

func TestGetUsers(t *testing.T) {
	t.Parallel()

	var (
		repository = repositories.NewRepositories(config.GetQuery())
		service    = services.NewServices(repository)
		h          = NewHandlers(service)
		echoMock   = mock.EchoMock{E: echo.New()}
		userId     = int(1)
	)

	config.NewTrunc().DeleteDataUsers()
	config.NewSeed().UsersSeeder()
	t.Run("success", func(t *testing.T) {
		c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
		c.SetPath("/jwt/v1/users")

		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.GetAllUsers(c)) {
			asserts.Equal(200, rec.Code)
		}
	})

	t.Run("failed", func(t *testing.T) {
		c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
		c.SetPath("/jwt/v1/users")
		c.SetParamNames("keywords")
		c.SetParamValues("test")

		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.GetAllUsers(c)) {
			asserts.Equal(404, rec.Code)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	t.Parallel()

	var (
		repository = repositories.NewRepositories(config.GetQuery())
		service    = services.NewServices(repository)
		h          = NewHandlers(service)
		echoMock   = mock.EchoMock{E: echo.New()}
		userId     = int(1)
	)

	config.NewTrunc().DeleteDataUsers()
	config.NewSeed().UsersSeeder()
	t.Run("success", func(t *testing.T) {
		c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
		c.SetPath("/jwt/v1/users:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.DeleteUser(c)) {
			asserts.Equal(200, rec.Code)
		}
	})

	t.Run("failed", func(t *testing.T) {
		c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
		c.SetPath("/jwt/v1/users:id")
		c.SetParamNames("id")
		c.SetParamValues("200")

		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.DeleteUser(c)) {
			asserts.Equal(404, rec.Code)
		}
	})
}
