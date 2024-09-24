package postgres

import (
	"github.com/hashicorp/go-uuid"
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

func (ar *albumRepository) GetAlbumId(trackId uint64) (uint64, error) {
	var track dao.TrackMeta

	tx := ar.db.Where("id = ?", trackId).Take(&track)
	if tx.Error != nil {
		return 0, errors.Wrap(tx.Error, "database error (table track)")
	}

	return track.AlbumID, nil
}

func (ar *albumRepository) IsAlbumOwned(albumId uint64, musicianId uint64) (bool, error) {
	var album dao.Album
	tx := ar.db.Where("id = ?", albumId).Take(&album)
	if tx.Error != nil {
		return false, errors.Wrap(tx.Error, "database error (table album)")
	}

	return musicianId == album.MusicianID, nil
}

func (ar *albumRepository) AddAlbumWithTracksOutbox(album *models.Album, tracks []*models.TrackMeta, musicianId uint64) (uint64, error) {
	pgAlbum := dao.ToPostgresAlbum(album, musicianId)
	var pgTracks []*dao.TrackMeta

	err := ar.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(pgAlbum).Error; err != nil {
			return err
		}

		for _, v := range tracks {
			var pgGenre dao.Genre
			txInner := tx.Where("name = ?", v.Genre).Take(&pgGenre)
			if txInner.Error != nil {
				return tx.Error
			}

			pgTracks = append(pgTracks, dao.ToPostgresTrack(v, pgGenre.ID, pgAlbum.ID))
		}

		if err := tx.Create(&pgTracks).Error; err != nil {
			return err
		}

		for _, v := range pgTracks {
			eventID, err := uuid.GenerateUUID()
			if err != nil {
				return err
			}

			var genre uint64
			genre = 0
			if v.GenreRefer != nil {
				genre = *v.GenreRefer
			}

			if err := tx.Create(&dao.Outbox{
				ID:         0,
				EventId:    eventID,
				TrackId:    v.ID,
				Source:     v.Source,
				Name:       v.Name,
				GenreRefer: genre,
				Type:       dao.TypeAdd,
				Sent:       false,
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "database error (table album)")
	}
	album.Id = pgAlbum.ID
	return pgAlbum.ID, nil
}

func (ar *albumRepository) DeleteAlbumOutbox(id uint64) error {
	err := ar.db.Transaction(func(tx *gorm.DB) error {
		var relations []*dao.TrackMeta
		if err := tx.Limit(dao.MaxLimit).Find(&relations, "album_id = ?", id).Error; err != nil {
			return err
		}

		for _, v := range relations {
			if err := tx.Delete(&dao.TrackMeta{}, v.ID).Error; err != nil {
				return err
			}

			eventID, err := uuid.GenerateUUID()
			if err != nil {
				return err
			}

			if err := tx.Create(&dao.Outbox{
				ID:         0,
				EventId:    eventID,
				TrackId:    v.ID,
				Source:     "",
				Name:       "",
				GenreRefer: 0,
				Type:       dao.TypeDelete,
				Sent:       false,
			}).Error; err != nil {
				return err
			}
		}

		if err := tx.Delete(&dao.Album{}, id).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "database error (table album)")
	}

	return nil
}
