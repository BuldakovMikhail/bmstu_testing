package repository

import "src/internal/models"

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type AlbumRepository interface {
	GetAlbum(id uint64) (*models.Album, error)
	AddAlbumWithTracksOutbox(album *models.Album, tracks []*models.TrackMeta, musicianId uint64) (uint64, error)
	DeleteAlbumOutbox(id uint64) error
	GetAllTracksForAlbum(albumId uint64) ([]*models.TrackMeta, error)

	IsAlbumOwned(albumId uint64, musicianId uint64) (bool, error)
	GetAlbumId(trackId uint64) (uint64, error)
}
