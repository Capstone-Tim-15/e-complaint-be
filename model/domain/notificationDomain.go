package domain

import "time"

type Notification struct {
	ID           string
	Complaint_ID string
	Complaint    Complaint
	Message      string
	Status       string
	CreatedAt    time.Time
}