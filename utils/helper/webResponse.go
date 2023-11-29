package helper

type TResponseMeta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TResponseMetaPage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	Total   int64  `json:"total"`
}

type TSuccessResponse struct {
	Meta    TResponseMeta `json:"meta"`
	Results interface{}   `json:"results"`
}

type TSuccessResponsePage struct {
	Meta    TResponseMetaPage `json:"meta"`
	Results interface{}       `json:"results"`
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

func SuccessResponsePage(message string, page int, limit int, total int64, data interface{}) interface{} {
	return TSuccessResponsePage{
		Meta: TResponseMetaPage{
			Success: true,
			Message: message,
			Page:    page,
			Limit:   limit,
			Total:   total,
		},
		Results: data,
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
