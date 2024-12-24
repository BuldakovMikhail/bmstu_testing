package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
	postgres2 "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"src/internal/domain/track/repository/postgres"
	"src/internal/lib/testing/builders"
	dbhelpers "src/internal/lib/testing/db"
	"src/internal/models"
	"src/internal/models/dao"
	"testing"
)

type TrackRepoSuite struct {
	suite.Suite
	t *testing.T
}

func (a *TrackRepoSuite) Test_GetTrack_Success(t provider.T) {
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

	t.Title("[GetTrack (repo)] Success")
	t.Tags("track")
	t.Parallel()
	t.WithNewStep("[GetTrack (repo)] Success", func(sCtx provider.StepCtx) {
		trackDao := builders.TrackDaoMetaBuilder{}.
			WithId(1).
			WithName("aboba").
			WithPayload([]byte{1, 2, 3}).
			WithAlbumId(1).
			Build()

		rows := dbhelpers.MapTracks([]*dao.Track{trackDao})

		mock.ExpectQuery("^SELECT (.+) FROM \"tracks\" WHERE id = (.+)$").
			WithArgs(1, 1).
			WillReturnRows(rows)

		expTrack := dao.ToModelTrackObject(trackDao)

		track, err := postgres.NewTrackRepository(gormDB).GetTrack(1)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(track)
		sCtx.Assert().Equal(expTrack, track)
	})
}

func (a *TrackRepoSuite) Test_GetTrack_Error(t provider.T) {
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

	t.Title("[GetTrack (repo)] Error from db")
	t.Tags("track, repo")
	t.Parallel()
	t.WithNewStep("[GetTrack (repo)] Error from db", func(sCtx provider.StepCtx) {
		mock.ExpectQuery("^SELECT (.+) FROM \"tracks\" WHERE id = (.+)$").
			WithArgs(1, 1).
			WillReturnError(assert.AnError)

		track, err := postgres.NewTrackRepository(gormDB).GetTrack(1)

		sCtx.Assert().ErrorIs(err, assert.AnError)
		sCtx.Assert().Nil(track)
	})
}

func (a *TrackRepoSuite) Test_GetTracksByPartName_Success(t provider.T) {
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

	t.Title("[GetTracksByPartName (repo)] Success")
	t.Tags("track, repo")
	t.Parallel()
	t.WithNewStep("[GetTracksByPartName (repo)] Success", func(sCtx provider.StepCtx) {
		trackDao1 := builders.TrackDaoMetaBuilder{}.
			WithId(1).
			WithName("aboba").
			WithPayload([]byte{1, 2, 3}).
			WithAlbumId(1).
			Build()

		trackDao2 := builders.TrackDaoMetaBuilder{}.
			WithId(2).
			WithName("zzzbobazzz").
			WithPayload([]byte{4, 5, 6}).
			WithAlbumId(1).
			Build()

		daoTracks := []*dao.Track{trackDao1, trackDao2}
		rows := dbhelpers.MapTracks(daoTracks)

		mock.ExpectQuery("^SELECT (.+) FROM \"tracks\" WHERE name LIKE(.+)$").
			WithArgs("%boba%", 2).
			WillReturnRows(rows)

		var expTracks []*models.TrackMeta
		for _, v := range daoTracks {
			expTracks = append(expTracks, dao.ToModelTrackMeta(v))
		}

		tracks, err := postgres.NewTrackRepository(gormDB).GetTracksByPartName("boba", 0, 2)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(tracks)
		sCtx.Assert().Equal(expTracks, tracks)
	})
}

func (a *TrackRepoSuite) Test_GetTracksByPartName_Error(t provider.T) {
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

	t.Title("[GetTracksByPartName (repo)] Error from db")
	t.Tags("track, repo")
	t.Parallel()
	t.WithNewStep("[GetTracksByPartName (repo)] Error from db", func(sCtx provider.StepCtx) {
		mock.ExpectQuery("^SELECT (.+) FROM \"tracks\" WHERE name LIKE(.+)$").
			WithArgs("%boba%", 2).
			WillReturnError(assert.AnError)

		tracks, err := postgres.NewTrackRepository(gormDB).GetTracksByPartName("boba", 0, 2)

		sCtx.Assert().ErrorIs(err, assert.AnError)
		sCtx.Assert().Nil(tracks)
	})
}
