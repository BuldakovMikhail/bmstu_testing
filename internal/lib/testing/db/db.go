package dbhelpers

import (
	"context"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"src/internal/lib/testhelpers"
	"src/internal/lib/testing/builders"
	"src/internal/lib/testing/mother"
	"src/internal/models/dao"
	"time"
)

type TestDatabaseMeta struct {
	DB        *gorm.DB
	Container *testhelpers.PostgresContainer
	IDs       map[string]uint64
}

func CreateDatabase(ctx context.Context, pathToInitScript string) (*TestDatabaseMeta, error) {
	ctx, cancel := context.WithTimeout(ctx, 240*time.Second)
	defer cancel()

	pgContainer, err := testhelpers.CreatePostgresContainer(ctx, pathToInitScript)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(pgContainer.ConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var ids = make(map[string]uint64)

	musicianId, err := initMusicianTable(db)
	if err != nil {
		pgContainer.Terminate(ctx)
		return nil, err
	}
	ids["musicianId"] = musicianId

	albumId, err := initAlbumTable(db, musicianId)
	if err != nil {
		pgContainer.Terminate(ctx)
		return nil, err
	}
	ids["albumId"] = albumId

	trackId, err := initTrackTable(db, albumId)
	if err != nil {
		pgContainer.Terminate(ctx)
		return nil, err
	}
	ids["trackId"] = trackId

	return &TestDatabaseMeta{
		DB:        db,
		Container: pgContainer,
		IDs:       ids,
	}, nil
}

func (t *TestDatabaseMeta) Terminate(ctx context.Context) {
	defer func() {
		err := t.Container.Terminate(ctx)
		if err != nil {
			log.Error().Err(err)
		}
	}()

	if err := t.DB.Delete(&dao.Track{}, t.IDs["trackId"]).Error; err != nil {
		log.Error().Err(err)
	}

	if err := t.DB.Delete(&dao.Album{}, t.IDs["albumId"]).Error; err != nil {
		log.Error().Err(err)
	}

	if err := t.DB.Delete(&dao.Musician{}, t.IDs["musicianId"]).Error; err != nil {
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
