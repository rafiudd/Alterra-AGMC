package controllers

import (
	"dynamic-crud/models"
	"dynamic-crud/transport"
	"dynamic-crud/utils"
)

func (c *controllers) CreateUser(user *models.User) (*transport.Response, error) {

	hash, _ := utils.HashPassword(user.Password)
	userMapping := &models.User{
		Password: hash,
		Name:     user.Name,
		Email:    user.Email,
	}

	err := c.repo.CreateUser(userMapping)
	if err != nil {
		return nil, err
	}

	result := &transport.Response{
		Code:    201,
		Status:  "success",
		Message: "Success create user",
		Data:    err,
	}

	return result, err
}

func (c *controllers) UpdateUser(user *models.User, id int) (*transport.Response, error) {
	userMapping := &models.User{
		Name:  user.Name,
		Email: user.Email,
	}

	err := c.repo.UpdateUser(userMapping, id)
	if err != nil {
		return nil, err
	}

	result := &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "Success update user",
		Data:    err,
	}

	return result, nil
}

func (c *controllers) DeleteUser(id int) (*transport.Response, error) {

	err := c.repo.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	result := &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "Success delete user",
		Data:    err,
	}

	return result, nil
}

func (c *controllers) GetUserById(id int) (*transport.Response, error) {

	user, err := c.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	userMapping := &transport.UserMapping{
		Name:      user.Name,
		Email:     user.Email,
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	result := &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "Success get user",
		Data:    userMapping,
	}
	return result, nil
}

func (c *controllers) GetAllUsers(keywords string) (*transport.Response, error) {
	users, err := c.repo.GetAllUsers(keywords)
	if err != nil {
		return nil, err
	}

	var arrUsers []transport.UserMapping
	for _, user := range users {
		userMapping := transport.UserMapping{
			Name:      user.Name,
			Email:     user.Email,
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		arrUsers = append(arrUsers, userMapping)
	}

	result := &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "Success get all users",
		Data:    arrUsers,
	}
	return result, nil
}
