package web

import "ecomplaint/model/schema"

type CategoryResponse struct {
	Id           string             `json:"id"`
	CategoryName string             `json:"CategoryName"`
	CreatedAt    string             `json:"created_at"`
	FAQ          []schema.FAQ       `json:"faq"`
	Complaint    []schema.Complaint `json:"complaint"`
}
