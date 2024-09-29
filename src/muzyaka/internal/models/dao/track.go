package dao

import "src/internal/models"

type Track struct {
	ID      uint64 `gorm:"column:id"`
	Payload []byte `gorm:"payload"`
	Name    string `gorm:"column:name"`
	AlbumID uint64 `gorm:"column:album_id"`
}

func (Track) TableName() string {
	return "tracks"
}

func ToPostgresTrack(e *models.TrackMeta, albumId uint64) *Track {
	return &Track{
		ID:      e.Id,
		Name:    e.Name,
		AlbumID: albumId,
	}
}

func ToModelTrackMeta(track *Track) *models.TrackMeta {
	return &models.TrackMeta{
		Id:   track.ID,
		Name: track.Name,
	}
}

func ToModelTrackObject(track *Track) *models.TrackObject {
	return &models.TrackObject{
		Payload: track.Payload,
	}
}
