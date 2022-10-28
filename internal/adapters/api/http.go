package api

import "zuri-stage-one/internal/ports"

type HTTPHandler struct {
	UserService ports.UserService
}

func NewHTTPHandler(userService ports.UserService) *HTTPHandler {
	return &HTTPHandler{
		UserService: userService,
	}
}
