package handlers

import (
	"DAY_4/config"
	"DAY_4/mock"
	"DAY_4/repositories"
	"DAY_4/services"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	var (
		echoMock   = mock.EchoMock{E: echo.New()}
		repository = repositories.NewRepositories(config.GetQuery())
		service    = services.NewServices(repository)
		h          = NewHandlers(service)
	)

	t.Run("success", func(t *testing.T) {

		c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
		c.SetPath("/v1/healthcheck")

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.HealthCheck(c)) {
			asserts.Equal(200, rec.Code)
		}
	})
}
