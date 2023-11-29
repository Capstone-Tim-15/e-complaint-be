package helper

type TResponseMeta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TSuccessResponse struct {
	Meta    TResponseMeta `json:"meta"`
	Results interface{}   `json:"results"`
}

type TErrorResponse struct {
	Meta TResponseMeta `json:"meta"`
}

func SuccessResponse(message string, data interface{}) interface{} {
	if data == nil {
		return TErrorResponse{
			Meta: TResponseMeta{
				Success: true,
				Message: message,
			},
		}
	} else {
		return TSuccessResponse{
			Meta: TResponseMeta{
				Success: true,
				Message: message,
			},
			Results: data,
		}
	}
}

func ErrorResponse(message string) interface{} {
	return TErrorResponse{
		Meta: TResponseMeta{
			Success: false,
			Message: message,
		},
	}
}

type TPSuccessResponse struct {
	Meta       TResponseMeta `json:"meta"`
	Results    interface{}   `json:"results"`
	Pagination interface{}   `json:"pagination"`
}

type TPagination struct {
	Offset int   `json:"offset"`
	Limit  int   `json:"limit"`
	Total  int64 `json:"total"`
}

func Pagination(offset int, limit int, total int64) TPagination {
	return TPagination{
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}
}

func PaginationResponse(message string, data interface{}, pagination interface{}) interface{} {
	if data == nil {
		return TErrorResponse{
			Meta: TResponseMeta{
				Success: true,
				Message: message,
			},
		}
	} else {
		return TPSuccessResponse{
			Meta: TResponseMeta{
				Success: true,
				Message: message,
			},
			Results:    data,
			Pagination: pagination,
		}
	}
}
