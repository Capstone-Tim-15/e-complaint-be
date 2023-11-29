package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
	"ecomplaint/utils/res"
	"fmt"
)

func (categoryService *CategoryServiceImpl) Pagination(offset, limit int) ([]domain.Category, *web.Pagination, error) {
	result, total, err := categoryService.CategoryRepository.Pagination(offset, limit)

	if total == 0 {
		return nil, nil, fmt.Errorf("category not found")
	}

	if err != nil {
		return nil, nil, fmt.Errorf("internal Server Error")
	}

	pagination := res.RecordToPaginationResponse(offset, limit, total)

	return result, pagination, nil
}
