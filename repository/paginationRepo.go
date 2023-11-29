package repository

import "ecomplaint/model/domain"

func (categoryRepository *CategoryRepositoryImpl) Pagination(offset int, limit int) ([]domain.Category, int64, error) {

	if offset < 0 || limit < 0 {
		return nil, 0, nil
	}

	category := []domain.Category{}
	var total int64

	query := categoryRepository.DB.Model(&category)

	query.Find(&category).Count(&total)

	query = query.Limit(limit).Offset(offset)

	result := query.Find(&category)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 {
		return nil, 0, nil
	}

	if offset >= int(total) {
		return nil, 0, nil
	}

	return category, total, nil
}
