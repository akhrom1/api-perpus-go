package services

import (
	"api-perpus-go/internal/models"
	"api-perpus-go/internal/repositories"
	"errors"
)

func CreateBook(title, author string, year, pages int, categoryID *int) (*models.Book, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}
	if author == "" {
		return nil, errors.New("author is required")
	}
	if year <= 0 {
		return nil, errors.New("year is required")
	}
	if pages <= 0 {
		return nil, errors.New("pages is required")
	}
	if categoryID == nil {
		return nil, errors.New("category is required")
	}

	exist, err := repositories.IsTitleExist(title, author)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("title already exists")
	}

	
	b := &models.Book{
		Title:      title,
		Author:     author,
		Year:		year,
		Pages: 		pages,
		CategoryID: categoryID,
		Status:     "AVAILABLE",
	}

	if err := repositories.CreateBook(b); err != nil {
		return nil, err
	}
	return b, nil
}

func ListBooks() ([]models.BookList, error) {
	return repositories.GetAllBooks()
}

func GetBook(id int) (*models.Book, error) {
	return repositories.GetBookByID(id)
}

func UpdateBook(id int, b *models.Book) error {
	if b.Title == "" {
		return errors.New("title is required")
	}
	return repositories.UpdateBook(id, b)
}

func DeleteBook(id int) error {
	book, err := repositories.GetBookByID(id)
	if err != nil {
		return err
	}
	if book.Status == "BORROWED" {
		return errors.New("book is currently borrowed")
	}
	return repositories.DeleteBook(id)
}
