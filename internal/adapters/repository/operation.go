package repository

import "zuri-stage-one/internal/core/models"

func (p *Postgres) Calculate(operation *models.Operation) (*models.Operation, error) {
	if err := p.DB.Create(operation).Error; err != nil {
		return nil, err
	}
	return operation, nil
}
