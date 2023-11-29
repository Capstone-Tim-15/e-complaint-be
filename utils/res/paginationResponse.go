package res

import "ecomplaint/model/web"

func RecordToPaginationResponse(offset, limit int, total int64) *web.Pagination {
	return &web.Pagination{
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}
}
