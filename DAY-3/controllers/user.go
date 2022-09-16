package controllers

import (
	m "DAY_3/middleware"
	"DAY_3/models"
	"DAY_3/transport"
	"DAY_3/utils"

	"golang.org/x/crypto/bcrypt"
)

func (c *controllers) CreateUser(user *models.User) (*transport.Response, error) {

	hash, _ := utils.HashPassword(user.Password)
	userMapping := &models.User{
		Password: hash,
		Username: user.Username,
		Email:    user.Email,
	}

	err := c.repo.CreateUser(userMapping)
	if err != nil {
		return nil, err
	}

	result := &transport.Response{
		Code:    201,
		Status:  "success",
		Message: "Success to create user",
		Data:    err,
	}

	return result, err
}

func (c *controllers) UpdateUser(user *models.User, id int) (*transport.Response, error) {
	userMapping := &models.User{
		Username: user.Username,
		Email:    user.Email,
	}

	err := c.repo.UpdateUser(userMapping, id)
	if err != nil {
		return nil, err
	}

	result := &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "Success to update user",
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
		Message: "Success to delete user",
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
		Username:  user.Username,
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
			Username:  user.Username,
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

func (c *controllers) UserLogin(username, password string) (*transport.Response, error) {
	user, err := c.repo.UserLogin(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// If the two passwords don't match, return a 401 status.
		return nil, err
	}

	userReplicated := transport.LoginResponse{}
	userReplicated.Email = user.Email
	userReplicated.Username = user.Username
	userReplicated.ID = user.ID
	userReplicated.Token, err = m.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	result := &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "Success Login",
		Data:    userReplicated,
	}
	return result, nil
}
