package repository

import "zuri-stage-one/internal/core/models"

func (p *Postgres) Get() ([]models.User, error) {
	var users []models.User

	if err := p.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
