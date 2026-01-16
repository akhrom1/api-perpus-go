package repositories

import (
	"api-perpus-go/config"
	"api-perpus-go/internal/models"
)

func CreateCategory(c *models.Category) error {
	query := `INSERT INTO categories (name) VALUES ($1) RETURNING id, created_at`
	return config.DB.QueryRow(query, c.Name).Scan(&c.ID, &c.CreatedAt)
}

func GetAllCategories() ([]models.Category, error) {
	rows, err := config.DB.Query(`SELECT id, name, created_at FROM categories ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func GetCategoryByID(id int) (*models.Category, error) {
	var c models.Category
	err := config.DB.QueryRow(
		`SELECT id, name, created_at FROM categories WHERE id=$1`, id,
	).Scan(&c.ID, &c.Name, &c.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func UpdateCategory(id int, name string) error {
	_, err := config.DB.Exec(`UPDATE categories SET name=$1 WHERE id=$2`, name, id)
	return err
}

func DeleteCategory(id int) error {
	_, err := config.DB.Exec(`DELETE FROM categories WHERE id=$1`, id)
	return err
}
