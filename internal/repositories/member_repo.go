package repositories

import (
	"api-perpus-go/config"
	"api-perpus-go/internal/models"
)

func CreateMember(m *models.Member) error {
	return config.DB.QueryRow(`
		INSERT INTO members (name, address,  phone)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`, m.Name, m.Address, m.Phone).
		Scan(&m.ID, &m.CreatedAt)
}

func GetAllMembers() ([]models.Member, error) {
	rows, err := config.DB.Query(`
		SELECT id, name, address, phone, created_at
		FROM members ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Member
	for rows.Next() {
		var m models.Member
		if err := rows.Scan(&m.ID, &m.Name, &m.Address, &m.Phone, &m.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	return list, nil
}

func GetMemberByID(id int) (*models.Member, error) {
	var m models.Member
	err := config.DB.QueryRow(`
		SELECT id, name, address, phone, created_at
		FROM members WHERE id=$1
	`, id).
		Scan(&m.ID, &m.Name, &m.Address, &m.Phone, &m.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &m, nil
}

func UpdateMember(id int, m *models.Member) error {
	_, err := config.DB.Exec(`
		UPDATE members SET name=$1, address=$2, class=$3,
		WHERE id=$4
	`, m.Name, m.Address, m.Phone, id)
	return err
}

func DeleteMember(id int) error {
	_, err := config.DB.Exec(`DELETE FROM members WHERE id=$1`, id)
	return err
}

func IsPhoneExist(phone string) (bool, error) {
	var count int
	err := config.DB.QueryRow(
		"SELECT COUNT(1) FROM members WHERE phone=$1",
		phone,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
