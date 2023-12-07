package web

type DashboardResponse struct {
	TotalUser      int64 `json:"total_user"`
	TotalComplaint int64 `json:"total_complaint"`
	TotalResolved  int64 `json:"total_resolved"`
}
