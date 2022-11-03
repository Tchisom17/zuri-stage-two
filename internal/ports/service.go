package ports

import "zuri-stage-one/internal/core/models"

type UserService interface {
	Get() (*models.User, error)
}

type OperationService interface {
	Calculate(operation *models.Operation) (*models.Operation, error)
}
