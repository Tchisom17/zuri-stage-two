package api

import "zuri-stage-one/internal/ports"

type HTTPHandler struct {
	UserService      ports.UserService
	OperationService ports.OperationService
}

func NewHTTPHandler(userService ports.UserService, operationService ports.OperationService) *HTTPHandler {
	return &HTTPHandler{
		UserService:      userService,
		OperationService: operationService,
	}
}
