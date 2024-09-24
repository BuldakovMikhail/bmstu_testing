package repository

import "src/internal/models"

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type AlbumRepository interface {
	AddAlbumWithTracksOutbox(album *models.Album, tracks []*models.TrackMeta, musicianId uint64) (uint64, error)
	DeleteAlbumOutbox(id uint64) error

	IsAlbumOwned(albumId uint64, musicianId uint64) (bool, error)
	GetAlbumId(trackId uint64) (uint64, error)
}
