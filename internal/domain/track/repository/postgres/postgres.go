package postgres

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"src/internal/domain/track/repository"
	"src/internal/models"
	"src/internal/models/dao"
)

type trackRepository struct {
	db *gorm.DB
}

func NewTrackRepository(db *gorm.DB) repository.TrackRepository {
	return &trackRepository{db: db}
}

func (t trackRepository) GetTracksByPartName(name string, offset int, limit int) ([]*models.TrackMeta, error) {
	var tracks []*dao.Track

	tx := t.db.
		Offset(offset).
		Limit(limit).
		Where("name LIKE ?", "%"+name+"%").
		Order("name").
		Find(&tracks)
	if err := tx.Error; err != nil {
		return nil, errors.Wrap(err, "database error (table track)")
	}

	var modelTracks []*models.TrackMeta

	for _, v := range tracks {
		modelTracks = append(modelTracks, dao.ToModelTrackMeta(v))
	}

	return modelTracks, nil
}

func (t trackRepository) GetTrack(id uint64) (*models.TrackObject, error) {
	var track dao.Track

	tx := t.db.Where("id = ?", id).Take(&track)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table track)")
	}

	return dao.ToModelTrackObject(&track), nil
}
