package models

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Complete    bool      `json:"complete"`
	CreatedAt   time.Time `json:"createdAt"`
}
