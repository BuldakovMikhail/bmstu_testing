package usecase

import (
	"github.com/pkg/errors"
	"src/internal/domain/album/repository"
	"src/internal/models"
)

type AlbumUseCase interface {
	GetAlbum(id uint64) (*models.Album, error)
	GetAllTracks(albumId uint64) ([]*models.TrackMeta, error)
}

type usecase struct {
	albumRep repository.AlbumRepository
}

var ErrAlbumIdIsZero = errors.New("album id cannot be zero")

func NewAlbumUseCase(albumRepository repository.AlbumRepository) AlbumUseCase {
	return &usecase{albumRep: albumRepository}
}

func (u *usecase) GetAlbum(id uint64) (*models.Album, error) {
	if id == 0 {
		return nil, ErrAlbumIdIsZero
	}

	res, err := u.albumRep.GetAlbum(id)

	if err != nil {
		return nil, errors.Wrap(err, "album.usecase.GetAlbum error while get")
	}

	return res, nil
}

func (u *usecase) GetAllTracks(albumId uint64) ([]*models.TrackMeta, error) {
	if albumId == 0 {
		return nil, ErrAlbumIdIsZero
	}

	tracks, err := u.albumRep.GetAllTracks(albumId)

	if err != nil {
		return nil, errors.Wrap(err, "album.usecase.GetAllTracks error while get")
	}

	return tracks, nil
}
