package web

type DashboardResponse struct {
	Monthly []Monthly `json:"month"`
}

type Monthly struct {
	Month string
	Total int64
}

type LandingPage struct {
	TotalUser      int64 `json:"total_user"`
	TotalComplaint int64 `json:"total_complaint"`
	TotalResolved  int64 `json:"total_resolved"`
}
