package ports

import "zuri-stage-one/internal/core/models"

type UserService interface {
	Get() (*models.User, error)
}
