package repository

import "src/internal/models"

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type UserRepository interface {
	GetUser(id uint64) (*models.User, error)
	AddUser(user *models.User) (uint64, error)
	GetUserByEmail(email string) (*models.User, error)

	AddUserWithMusician(musician *models.Musician, user *models.User) (uint64, error)
}
