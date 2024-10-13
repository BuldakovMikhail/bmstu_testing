package postgres

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"src/internal/domain/album/repository"
	"src/internal/models"
	"src/internal/models/dao"
)

type albumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) repository.AlbumRepository {
	return &albumRepository{db: db}
}

func (ar *albumRepository) GetAllTracks(albumId uint64) ([]*models.TrackMeta, error) {
	var tempTracks []*dao.Track

	tx := ar.db.Limit(dao.MaxLimit).Find(&tempTracks, "album_id = ?", albumId)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table album)")
	}

	if len(tempTracks) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var tracks []*models.TrackMeta
	for _, v := range tempTracks {
		tracks = append(tracks, dao.ToModelTrackMeta(v))
	}
	return tracks, nil
}

func (ar *albumRepository) GetAlbum(id uint64) (*models.Album, error) {
	var album dao.Album

	tx := ar.db.Where("id = ?", id).Take(&album)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table album)")
	}
	return dao.ToModelAlbum(&album), nil
}
