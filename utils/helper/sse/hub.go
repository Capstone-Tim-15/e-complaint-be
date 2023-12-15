package sse

import "ecomplaint/model/domain"

type Hub struct {
	ComplaintChannel map[string]chan domain.Complaint
}