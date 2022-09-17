package handlers

import (
	m "DAY_4/middleware"
	"DAY_4/models"
	"DAY_4/transport"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) CreateUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)

	if err := c.Validate(user); err != nil {
		return err
	}

	response := new(transport.Response)
	result, err := h.service.CreateUser(user)

	if err != nil {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to create user"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Data = result.Data
	return c.JSON(http.StatusCreated, response)
}

func (h *handler) UpdateUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(transport.Response)
	extractToken := m.ExtractTokenUserId(c)
	result, err := h.service.UpdateUser(user, idInt)

	if err != nil || float64(idInt) != extractToken {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Failed to update user"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Data = result.Data
	return c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(transport.Response)

	extractToken := m.ExtractTokenUserId(c)

	result, err := h.service.DeleteUser(idInt)

	if err != nil || float64(idInt) != extractToken {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Failed to delete user"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Data = result.Data

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetUserById(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(transport.Response)
	result, err := h.service.GetUserById(idInt)

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "User not found"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Data = result.Data

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetAllUsers(c echo.Context) error {
	response := new(transport.Response)
	result, err := h.service.GetAllUsers(c.QueryParam("keywords"))

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Users not found"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Data = result.Data

	return c.JSON(http.StatusOK, response)
}

func (h *handler) UserLogin(c echo.Context) error {
	response := new(transport.Response)
	user := new(models.User)
	c.Bind(user)
	result, err := h.service.UserLogin(user.Username, user.Password)

	if err != nil {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Login failed"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Data = result.Data

	return c.JSON(http.StatusOK, response)
}
