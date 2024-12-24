package repository

import "src/internal/models"

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type TrackRepository interface {
	GetTrack(id uint64) (*models.TrackObject, error)
	GetTracksByPartName(name string, offset int, limit int) ([]*models.TrackMeta, error)
}
