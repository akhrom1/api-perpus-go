package repositories

import (
	"api-perpus-go/config"
	"api-perpus-go/internal/models"
	"fmt"
	"time"
)

func CreateLoan(l *models.Loan) error {
	return config.DB.QueryRow(`
		INSERT INTO loans (member_id, loan_date, due_date, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, l.MemberID, l.LoanDate, l.DueDate, l.Status).
		Scan(&l.ID)
}

func AddLoanItem(loanID, bookID int) error {
	_, err := config.DB.Exec(`
		INSERT INTO loan_items (loan_id, book_id)
		VALUES ($1, $2)
	`, loanID, bookID)
	return err
}

func ReturnLoan(id int) error {
	_, err := config.DB.Exec(`
		UPDATE loans
		SET return_date = NOW(), status = 'RETURNED'
		WHERE id = $1
	`, id)
	return err
}

func GetLoanBooks(loanID int) ([]int, error) {
	rows, err := config.DB.Query(`SELECT book_id FROM loan_items WHERE loan_id=$1`, loanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		rows.Scan(&id)
		ids = append(ids, id)
	}
	return ids, nil
}

func CalculateFine(dueDate time.Time, returnDate *time.Time) float64 {
    var endDate time.Time
	//  hardcodedNow := time.Date(2026, 1, 25, 0, 0, 0, 0, time.UTC)
    if returnDate != nil {
        endDate = *returnDate
    } else {
        // endDate = hardcodedNow
        endDate = time.Now() 
    }

    overdueDays := int(endDate.Sub(dueDate).Hours() / 24)
    if overdueDays > 0 {
        return float64(overdueDays) * 5000
    }
    return 0
}



func GetAllLoans() ([]models.Loan, error) {
	rows, err := config.DB.Query(`
		SELECT id, member_id, loan_date, due_date, status
		FROM loans
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var loans []models.Loan
	for rows.Next() {
		var l models.Loan
		if err := rows.Scan(
			&l.ID,
			&l.MemberID,
			&l.LoanDate,
			&l.DueDate,
			&l.Status,
		); err != nil {
			return nil, err
		}

		fmt.Printf("LoanID: %d, MemberID: %d, DueDate: %v, ReturnDate: %v, Status: %s\n",
			l.ID, l.MemberID, l.DueDate, l.ReturnDate, l.Status)

		l.Fine = CalculateFine(l.DueDate, l.ReturnDate)

		fmt.Printf("LoanID: %d, Fine: %.2f\n", l.ID, l.Fine)
		loans = append(loans, l)
	}

	return loans, nil
}

func IsBookAvailable(bookID int) (bool, error) {
	var status string
	err := config.DB.QueryRow(
		"SELECT status FROM books WHERE id=$1",
		bookID,
	).Scan(&status)

	if err != nil {
		return false, err
	}

	return status == "AVAILABLE", nil
}

func GetLoanItemsByLoanID(loanID int) ([]models.LoanItem, error) {
	rows, err := config.DB.Query(`
		SELECT li.id, li.loan_id, li.book_id, b.title
		FROM loan_items li
		JOIN books b ON li.book_id = b.id
		WHERE li.loan_id=$1
	`, loanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.LoanItem
	for rows.Next() {
		var li models.LoanItem
		if err := rows.Scan(&li.ID, &li.LoanID, &li.BookID, &li.Title); err != nil {
			return nil, err
		}
		items = append(items, li)
	}
	return items, nil
}


func CountLoanItems(loanID int) (int, error) {
	var count int
	err := config.DB.QueryRow(`
		SELECT COUNT(*) FROM loan_items WHERE loan_id=$1
	`, loanID).Scan(&count)
	return count, err
}



func GetLoanByID(id int) (*models.Loan, error) {
	var loan models.Loan

	err := config.DB.QueryRow(`
		SELECT id, member_id, loan_date, due_date, return_date, status
		FROM loans
		WHERE id=$1
	`, id).Scan(&loan.ID, &loan.MemberID, &loan.LoanDate, &loan.DueDate, &loan.ReturnDate, &loan.Status)
	if err != nil {
		return nil, err
	}

	return &loan, nil
}
