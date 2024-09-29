package repository

import "src/internal/models"

//go:generate mockgen -source=repository.go -destination=mocks/mock.go
type AlbumRepository interface {
	GetAlbum(id uint64) (*models.Album, error)
	GetAllTracks(albumId uint64) ([]*models.TrackMeta, error)
}
