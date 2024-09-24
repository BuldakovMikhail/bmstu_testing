package postgres

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	repository2 "src/internal/domain/user/repository"
	"src/internal/models"
	"src/internal/models/dao"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository2.UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user dao.User
	tx := u.db.Where("email = ?", email).Take(&user)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table users_musicians)")
	}

	return dao.ToModelUser(&user), nil
}

func (u userRepository) GetUser(id uint64) (*models.User, error) {
	var user dao.User

	tx := u.db.Where("id = ?", id).Take(&user)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table user)")
	}

	return dao.ToModelUser(&user), nil
}

func (u userRepository) AddUserWithMusician(musician *models.Musician, user *models.User) (uint64, error) {
	pgMusician := dao.ToPostgresMusician(musician)
	pgUser := dao.ToPostgresUser(user)

	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&pgMusician).Error; err != nil {
			return err
		}

		temp := musician
		temp.Id = pgMusician.ID
		pgMusicianPhotos := dao.ToPostgresMusicianPhotos(temp)
		if len(pgMusicianPhotos) != 0 {
			if err := tx.Create(&pgMusicianPhotos).Error; err != nil {
				return err
			}
		}

		if err := tx.Create(&pgUser).Error; err != nil {
			return err
		}

		pgRelation := dao.UserMusician{
			UserId:     pgUser.ID,
			MusicianId: pgMusician.ID,
		}

		if err := tx.Create(&pgRelation).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "database error (table musician)")
	}
	musician.Id = pgMusician.ID
	user.Id = pgUser.ID
	return pgUser.ID, nil
}

func (u userRepository) AddUser(user *models.User) (uint64, error) {
	pgUser := dao.ToPostgresUser(user)

	tx := u.db.Create(&pgUser)
	if tx.Error != nil {
		return 0, errors.Wrap(tx.Error, "database error (table user)")
	}

	user.Id = pgUser.ID
	return pgUser.ID, nil
}
