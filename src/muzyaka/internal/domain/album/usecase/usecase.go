package usecase

import (
	"src/internal/domain/album/repository"
	repository2 "src/internal/domain/track/repository"
	"src/internal/models"
)

type AlbumUseCase interface {
	GetAlbum(id uint64) (*models.Album, error)
	GetAllTracks(albumId uint64) ([]*models.TrackMeta, error)
}

type usecase struct {
	albumRep repository.AlbumRepository
	trackRep repository2.TrackRepository
}

func NewAlbumUseCase(albumRepository repository.AlbumRepository,
	trackRepository repository2.TrackRepository) AlbumUseCase {
	return &usecase{albumRep: albumRepository, trackRep: trackRepository}
}
