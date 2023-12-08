package web

type DashboardResponse struct {
	Monthly []Monthly `json:"month"`
}

type Monthly struct {
	Month string
	Total int64
}
