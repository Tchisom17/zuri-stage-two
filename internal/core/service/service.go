package service

import (
	"strings"
	"zuri-stage-one/internal/core/helpers"
	"zuri-stage-one/internal/core/models"
	"zuri-stage-one/internal/ports"
)

type userService struct {
	userRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) ports.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

type operationService struct {
	operationRepository ports.OperationRepository
}

func NewOperationService(operationRepository ports.OperationRepository) ports.OperationService {
	return &operationService{
		operationRepository: operationRepository,
	}
}

func (u *userService) Get() (*models.User, error) {
	return u.userRepository.Get()
}

func (o *operationService) Calculate(operation *models.Operation) (*models.Operation, error) {
	result := operation.ReadString(operation.OperationType)
	if operation.X == 0 && operation.Y == 0 {
		operation.X = result[0]
		operation.Y = result[len(result)-1]
	}

	if strings.Contains(operation.OperationType, "add") || strings.Contains(operation.OperationType, "plus") || strings.Contains(operation.OperationType, "sum") {
		operation.OperationType = "addition"
	} else if strings.Contains(operation.OperationType, "subtract") || strings.Contains(operation.OperationType, "minus") || strings.Contains(operation.OperationType, "take away") {
		operation.OperationType = "subtraction"
	} else if strings.Contains(operation.OperationType, "times") || strings.Contains(operation.OperationType, "multiply") || strings.Contains(operation.OperationType, "product") {
		operation.OperationType = "multiplication"
	}

	_, err := helpers.OperationEnum(operation.OperationType)
	if err != nil {
		return nil, err
	}

	if operation.OperationType == "addition" {
		operation.Result = operation.X + operation.Y
	} else if operation.OperationType == "subtraction" {
		operation.Result = operation.X - operation.Y
	} else {
		operation.Result = operation.X * operation.Y
	}

	return o.operationRepository.Calculate(operation)
}
