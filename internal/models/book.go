package models

import "time"

type Book struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Year       int       `json:"year"`
	Pages      int       `json:"pages"`
	CategoryID *int      `json:"category_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

type BookList struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
