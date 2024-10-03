package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
	postgres2 "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"src/internal/domain/album/repository/postgres"
	"src/internal/lib/testing/builders"
	dbhelpers "src/internal/lib/testing/db"
	"src/internal/models"
	"src/internal/models/dao"
	"testing"
)

type AlbumRepoSuite struct {
	suite.Suite
	t *testing.T
}

func (a *AlbumRepoSuite) Test_GetAlbum_Success(t provider.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		a.t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres2.New(postgres2.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		a.t.Fatalf("an error '%s' was not expected when creating gormDB", err)
	}

	t.Title("[GetAlbum (repo)] Success")
	t.Tags("album, repo")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		albumDao := builders.AlbumDaoBuilder{}.
			WithId(1).
			WithName("test").
			WithCoverFile([]byte{1, 2, 3}).
			WithMusicianId(1).
			Build()

		rows := dbhelpers.MapAlbum(albumDao)

		mock.ExpectQuery("^SELECT (.+) FROM \"albums\" WHERE id = (.+)$").
			WithArgs(1, 1).
			WillReturnRows(rows)

		expAlbum := dao.ToModelAlbum(albumDao)

		album, err := postgres.NewAlbumRepository(gormDB).GetAlbum(1)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(album)
		sCtx.Assert().Equal(expAlbum, album)
	})
}

func (a *AlbumRepoSuite) Test_GetAlbum_Error(t provider.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		a.t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres2.New(postgres2.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		a.t.Fatalf("an error '%s' was not expected when creating gormDB", err)
	}

	t.Title("[GetAlbum (repo)] Error from db")
	t.Tags("album, repo")
	t.Parallel()
	t.WithNewStep("Error from db", func(sCtx provider.StepCtx) {
		mock.ExpectQuery("^SELECT (.+) FROM \"albums\" WHERE id = (.+)$").
			WithArgs(1, 1).
			WillReturnError(assert.AnError)

		album, err := postgres.NewAlbumRepository(gormDB).GetAlbum(1)

		sCtx.Assert().ErrorIs(err, assert.AnError)
		sCtx.Assert().Nil(album)
	})
}

func (a *AlbumRepoSuite) Test_GetAllTracks_Success(t provider.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		a.t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres2.New(postgres2.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		a.t.Fatalf("an error '%s' was not expected when creating gormDB", err)
	}

	t.Title("[GetAllTracks (repo)] Success")
	t.Tags("album, repo")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		trackDao := builders.TrackDaoMetaBuilder{}.
			WithId(1).
			WithName("a").
			WithPayload([]byte{1, 2, 3}).
			WithAlbumId(1).
			Build()

		rows := dbhelpers.MapTracks([]*dao.Track{trackDao})

		mock.ExpectQuery("^SELECT (.+) FROM \"tracks\" WHERE album_id = (.+)$").
			WithArgs(1, 100).
			WillReturnRows(rows)

		expTracks := []*models.TrackMeta{dao.ToModelTrackMeta(trackDao)}

		tracks, err := postgres.NewAlbumRepository(gormDB).GetAllTracks(1)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(tracks)
		sCtx.Assert().Equal(expTracks, tracks)
	})
}

func (a *AlbumRepoSuite) Test_GetAllTracks_Error(t provider.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		a.t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres2.New(postgres2.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		a.t.Fatalf("an error '%s' was not expected when creating gormDB", err)
	}

	t.Title("[GetAllTracks (repo)] Error from db")
	t.Tags("album, repo")
	t.Parallel()
	t.WithNewStep("Error from db", func(sCtx provider.StepCtx) {
		mock.ExpectQuery("^SELECT (.+) FROM \"tracks\" WHERE album_id = (.+)$").
			WithArgs(1, 100).
			WillReturnError(assert.AnError)

		tracks, err := postgres.NewAlbumRepository(gormDB).GetAllTracks(1)

		sCtx.Assert().ErrorIs(err, assert.AnError)
		sCtx.Assert().Nil(tracks)
	})
}
