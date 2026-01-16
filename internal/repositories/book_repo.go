package repositories

import (
	"api-perpus-go/config"
	"api-perpus-go/internal/models"
	"database/sql"
)

func CreateBook(b *models.Book) error {
	query := `
		INSERT INTO books (title, author, year, pages, category_id, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`
	return config.DB.QueryRow(
		query, b.Title, b.Author, b.Year, b.Pages, b.CategoryID, b.Status,
	).Scan(&b.ID, &b.CreatedAt)
}

func GetAllBooks() ([]models.BookList, error) {
	rows, err := config.DB.Query(`
		SELECT id, title, author, year
		FROM books
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.BookList

	for rows.Next() {
		var b models.BookList
		if err := rows.Scan(&b.ID, &b.Title, &b.Author,  &b.Year); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}


func GetBookByID(id int) (*models.Book, error) {
	var b models.Book
	var catID sql.NullInt64

	err := config.DB.QueryRow(`
		SELECT id, title, author, year, pages, category_id, status, created_at
		FROM books WHERE id=$1
	`, id).Scan(&b.ID, &b.Title, &b.Author, &b.Year, &b.Pages, &catID, &b.Status, &b.CreatedAt)

	if err != nil {
		return nil, err
	}

	if catID.Valid {
		v := int(catID.Int64)
		b.CategoryID = &v
	}

	return &b, nil
}

func UpdateBook(id int, b *models.Book) error {
	_, err := config.DB.Exec(`
		UPDATE books SET title=$1, author=$2, category_id=$3 WHERE id=$4
	`, b.Title, b.Author, b.CategoryID, id)
	return err
}

func DeleteBook(id int) error {
	_, err := config.DB.Exec(`DELETE FROM books WHERE id=$1`, id)
	return err
}

func UpdateBookStatus(id int, status string) error {
	_, err := config.DB.Exec(`UPDATE books SET status=$1 WHERE id=$2`, status, id)
	return err
}

func IsTitleExist(title, author string) (bool, error) {
	var count int
	err := config.DB.QueryRow(
		"SELECT COUNT(1) FROM Books WHERE title=$1 AND author=$2",
		title, author,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

