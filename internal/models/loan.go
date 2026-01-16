package models

import "time"

type Loan struct {
	ID         int        `json:"id"`
	MemberID  int        `json:"member_id"`
	LoanDate  time.Time  `json:"loan_date"`
	DueDate   time.Time  `json:"due_date"`
	ReturnDate *time.Time `json:"return_date"`
	Status    string     `json:"status"`
	Quantity   int         `json:"quantity"` 
	LoanItems   []LoanItem  `json:"loan_items,omitempty"`
	Fine       float64     `json:"fine"` 
}
