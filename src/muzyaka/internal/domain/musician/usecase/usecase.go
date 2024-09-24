package usecase

import (
	"github.com/pkg/errors"
	"src/internal/domain/musician/repository"
	"src/internal/models"
)

type MusicianUseCase interface {
	GetMusician(id uint64) (*models.Musician, error)

	GetMusicianIdForUser(userId uint64) (uint64, error)
}

type usecase struct {
	musicianRep repository.MusicianRepository
}

func NewMusicianUseCase(rep repository.MusicianRepository) MusicianUseCase {
	return &usecase{musicianRep: rep}
}

func (u *usecase) GetMusicianIdForUser(userId uint64) (uint64, error) {
	id, err := u.musicianRep.GetMusicianIdForUser(userId)
	if err != nil {
		return 0, errors.Wrap(err, "musician.usecase.GetMusicianIdForUser error while get")
	}

	return id, nil
}

func (u *usecase) GetMusician(id uint64) (*models.Musician, error) {
	res, err := u.musicianRep.GetMusician(id)

	if err != nil {
		return nil, errors.Wrap(err, "musician.usecase.GetMusician error while get")
	}

	return res, nil
}
