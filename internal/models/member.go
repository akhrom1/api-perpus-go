package models

import "time"

type Member struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Address    string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}
