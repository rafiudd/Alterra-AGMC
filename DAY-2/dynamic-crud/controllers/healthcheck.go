package controllers

import (
	"dynamic-crud/transport"
)

func (c *controllers) HealthCheck() *transport.Response {
	return &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "OK",
	}
}
