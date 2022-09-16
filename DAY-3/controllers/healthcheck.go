package controllers

import (
	"DAY_3/transport"
)

func (c *controllers) HealthCheck() *transport.Response {
	return &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "Server successfully running",
	}
}
