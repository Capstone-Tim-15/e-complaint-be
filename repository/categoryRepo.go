package repository

import "ecomplaint/model/domain"

func (repository *CategoryRepositoryImpl) FindAll() ([]domain.Category, error) {
	categories := []domain.Category{}

	result := repository.DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}
