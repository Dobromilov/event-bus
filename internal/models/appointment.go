package models

import "time"

type Appointment struct {
	ID        int64     `json:"id"`
	Client    string    `json:"client"`
	Service   string    `json:"service"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
