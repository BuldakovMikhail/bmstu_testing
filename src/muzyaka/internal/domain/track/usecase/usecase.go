package usecase

import (
	"github.com/pkg/errors"
	"src/internal/domain/track/repository"
	"src/internal/models"
)

const MinPageSize = 10
const MaxPageSize = 100

type TrackUseCase interface {
	GetTrack(id uint64) (*models.TrackObject, error)
	GetTracksByPartName(name string, page int, pageSize int) ([]*models.TrackMeta, error)
}

type usecase struct {
	trackRep   repository.TrackRepository
	storageRep repository.TrackStorage
}

func NewTrackUseCase(rep repository.TrackRepository, storage repository.TrackStorage) TrackUseCase {
	return &usecase{trackRep: rep, storageRep: storage}
}

func (u *usecase) GetTracksByPartName(name string, page int, pageSize int) ([]*models.TrackMeta, error) {
	if page <= 0 {
		page = 1
	}

	switch {
	case pageSize > MaxPageSize:
		pageSize = MaxPageSize
	case pageSize < MinPageSize:
		pageSize = MinPageSize
	}

	offset := (page - 1) * pageSize
	tracks, err := u.trackRep.GetTracksByPartName(name, offset, pageSize)
	if err != nil {
		return nil, errors.Wrap(err, "track.usecase.GetTracksByPartName error while get")
	}

	return tracks, nil
}

func (u *usecase) GetTrack(id uint64) (*models.TrackObject, error) {
	meta, err := u.trackRep.GetTrack(id)
	if err != nil {
		return nil, errors.Wrap(err, "track.usecase.GetTrack error while get")
	}

	res, err := u.storageRep.LoadObject(meta)
	if err != nil {
		return nil, errors.Wrap(err, "track.usecase.GetTrack error while get")
	}

	return res, nil
}
