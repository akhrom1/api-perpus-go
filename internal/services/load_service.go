package services

import (
	"errors"
	"time"

	"api-perpus-go/internal/models"
	"api-perpus-go/internal/repositories"
)

func CreateLoan(memberID int, bookIDs []int, dueDate time.Time) (*models.Loan, error) {
	if memberID == 0 || len(bookIDs) == 0 {
		return nil, errors.New("member and books required")
	}
	
	for _, bookID := range bookIDs {
		available, err := repositories.IsBookAvailable(bookID)
		if err != nil {
			return nil, err
		}
		if !available {
			return nil, errors.New("one or more books already borrowed")
		}
	}


	loanDate := time.Now()
	if dueDate.IsZero() {
		dueDate = loanDate.AddDate(0, 0, 7)
	}

	loan := &models.Loan{
		MemberID: memberID,
		LoanDate: time.Now(),
		DueDate:  dueDate,
		Status:   "BORROWED",
	}

	if err := repositories.CreateLoan(loan); err != nil {
		return nil, err
	}

	for _, bookID := range bookIDs {
		if err := repositories.UpdateBookStatus(bookID, "BORROWED"); err != nil {
			return nil, err
		}
		if err := repositories.AddLoanItem(loan.ID, bookID); err != nil {
			return nil, err
		}
	}

	return loan, nil
	
	
	
}

func ReturnLoan(id int) error {
	books, err := repositories.GetLoanBooks(id)
	if err != nil {
		return err
	}

	for _, bookID := range books {
		repositories.UpdateBookStatus(bookID, "AVAILABLE")
	}

	return repositories.ReturnLoan(id)
}


func GetAllLoans() ([]models.Loan, error) {
	loans, err := repositories.GetAllLoans()
	if err != nil {
		return nil, err
	}

	for i := range loans {
		count, err := repositories.CountLoanItems(loans[i].ID)
		if err != nil {
			return nil, err
		}
		loans[i].Quantity = count
	}

	return loans, nil
}


func GetLoanByID(id int) (*models.Loan, error) {
	loan, err := repositories.GetLoanByID(id)
	if err != nil {
		return nil, err
	}

	items, err := repositories.GetLoanItemsByLoanID(loan.ID)
	if err != nil {
		return nil, err
	}
	loan.LoanItems = items
	loan.Quantity = len(items)

	return loan, nil
}

