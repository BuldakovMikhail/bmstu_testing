package e2e

import (
	"context"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"src/internal/lib/testing/builders"
	"src/internal/lib/testing/mother"
	"src/internal/models/dao"
)

const dsn = "host=db user=postgres password=123 dbname=postgres port=5432"

func InitDatabase(ctx context.Context) (*gorm.DB, map[string]uint64, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	var ids = make(map[string]uint64)

	musicianId, err := initMusicianTable(db)
	if err != nil {
		return nil, nil, err
	}
	ids["musicianId"] = musicianId

	albumId, err := initAlbumTable(db, musicianId)
	if err != nil {
		return nil, nil, err
	}
	ids["albumId"] = albumId

	trackId, err := initTrackTable(db, albumId)
	if err != nil {
		return nil, nil, err
	}
	ids["trackId"] = trackId

	return db, ids, nil
}

func ClearTestDB(db *gorm.DB, ids map[string]uint64) {
	if err := db.Delete(&dao.Track{}, ids["trackId"]).Error; err != nil {
		log.Error().Err(err)
	}

	if err := db.Delete(&dao.Album{}, ids["albumId"]).Error; err != nil {
		log.Error().Err(err)
	}

	if err := db.Delete(&dao.Musician{}, ids["musicianId"]).Error; err != nil {
		log.Error().Err(err)
	}
}

func initMusicianTable(db *gorm.DB) (uint64, error) {
	musician := mother.MusicianDaoMother{}.DefaultMusician()
	if err := db.Create(musician).Error; err != nil {
		return 0, err
	}

	return musician.ID, nil
}

func initAlbumTable(db *gorm.DB, musicianId uint64) (uint64, error) {
	album := builders.AlbumDaoBuilder{}.
		WithMusicianId(musicianId).
		WithName("test").
		WithCoverFile([]byte{1, 2, 3}).
		Build()

	if err := db.Create(album).Error; err != nil {
		return 0, nil
	}

	return album.ID, nil
}

func initTrackTable(db *gorm.DB, albumId uint64) (uint64, error) {
	track := builders.TrackDaoMetaBuilder{}.
		WithName("test").
		WithPayload([]byte{4, 5, 6}).
		WithAlbumId(albumId).
		Build()

	if err := db.Create(track).Error; err != nil {
		return 0, nil
	}

	return track.ID, nil
}
