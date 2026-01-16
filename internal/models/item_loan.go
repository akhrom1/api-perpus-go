package models

type LoanItem struct {
	ID     int `json:"id"`
	LoanID int `json:"loan_id"`
	BookID int `json:"book_id"`
	Title  string `json:"title"`
}
