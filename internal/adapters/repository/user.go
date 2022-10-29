package repository

import "zuri-stage-one/internal/core/models"

func (p *Postgres) Get() (*models.User, error) {
	user := &models.User{}

	if err := p.DB.Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
