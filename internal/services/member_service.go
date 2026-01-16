package services

import (
	"api-perpus-go/internal/models"
	"api-perpus-go/internal/repositories"
	"errors"
)

func CreateMember(name, address, phone string) (*models.Member, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if address == "" {
		return nil, errors.New("address is required")
	}
	if phone == "" {
		return nil, errors.New("phone is required")
	}

	exist, err := repositories.IsPhoneExist(phone)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("phone already exists")
	}

	m := &models.Member{
		Name:   name,
		Address: address,
		Phone:  phone,
	}

	if err := repositories.CreateMember(m); err != nil {
		return nil, err
	}
	return m, nil
}

func ListMembers() ([]models.Member, error) {
	return repositories.GetAllMembers()
}

func GetMember(id int) (*models.Member, error) {
	return repositories.GetMemberByID(id)
}

func UpdateMember(id int, m *models.Member) error {
	if m.Name == "" || m.Phone == "" {
		return errors.New("name and phone are required")
	}
	return repositories.UpdateMember(id, m)
}

func DeleteMember(id int) error {
	return repositories.DeleteMember(id)
}
