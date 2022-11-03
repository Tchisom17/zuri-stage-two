package ports

import "zuri-stage-one/internal/core/models"

type UserRepository interface {
	Get() (*models.User, error)
}

type OperationRepository interface {
	Calculate(operation *models.Operation) (*models.Operation, error)
}
