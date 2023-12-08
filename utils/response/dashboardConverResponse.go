package response

import (
	"ecomplaint/model/web"
)

type DashboardComplaint struct {
	Dashboard []web.Monthly `json:"complaint_statistic"`
}
type DashboardUser struct {
	Dashboard []web.Monthly `json:"users_statistic"`
}
type DashboardSolved struct {
	Dashboard []web.Monthly `json:"solved_statistic"`
}

func DashComplaintResponse(req []web.Monthly) DashboardComplaint {
	var dashResponse DashboardComplaint
	for _, month := range req {
		dashResponse.Dashboard = append(dashResponse.Dashboard, month)
	}
	return dashResponse
}

func DashUserResponse(req []web.Monthly) DashboardUser {
	var dashResponse DashboardUser
	for _, month := range req {
		dashResponse.Dashboard = append(dashResponse.Dashboard, month)
	}
	return dashResponse
}

func DashSolvedResponse(req []web.Monthly) DashboardSolved {
	var dashResponse DashboardSolved
	for _, month := range req {
		dashResponse.Dashboard = append(dashResponse.Dashboard, month)
	}
	return dashResponse
}
