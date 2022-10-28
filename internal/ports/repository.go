package ports

import "zuri-stage-one/internal/core/models"

type UserRepository interface {
	Get() ([]models.User, error)
}
