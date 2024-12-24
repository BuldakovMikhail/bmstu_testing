package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
	postgres2 "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"src/internal/domain/track/repository/postgres"
	"src/internal/domain/track/usecase"
	"src/internal/lib/testing/builders"
	dbhelpers "src/internal/lib/testing/db"
	"src/internal/models"
	"src/internal/models/dao"
	"testing"
)

type TrackSuite struct {
	suite.Suite
	t *testing.T
}

func (a *TrackSuite) Test_GetTrack_Success(t provider.T) {
	c := gomock.NewController(t)
	defer c.Finish()

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

	t.Title("[GetTrack (service)] Success")
	t.Tags("track, service")
	t.Parallel()
	t.WithNewStep("[GetTrack (service)] Success", func(sCtx provider.StepCtx) {
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
		repo := postgres.NewTrackRepository(gormDB)

		track, err := usecase.NewTrackUseCase(repo).GetTrack(1)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(track)
		sCtx.Assert().Equal(expTrack, track)
	})
}

func (a *TrackSuite) Test_GetTrack_Error(t provider.T) {
	c := gomock.NewController(t)
	defer c.Finish()

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

	t.Title("[GetTrack (service)] Error from db")
	t.Tags("track, service")
	t.Parallel()
	t.WithNewStep("[GetTrack (service)] Error from db", func(sCtx provider.StepCtx) {
		mock.ExpectQuery("^SELECT (.+) FROM \"tracks\" WHERE id = (.+)$").
			WithArgs(1, 1).
			WillReturnError(assert.AnError)

		repo := postgres.NewTrackRepository(gormDB)

		track, err := usecase.NewTrackUseCase(repo).GetTrack(1)

		sCtx.Assert().ErrorIs(err, assert.AnError)
		sCtx.Assert().Nil(track)
	})
}

func (a *TrackSuite) Test_GetTracksByPartName_Success(t provider.T) {
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

	t.Title("[GetTracksByPartName (service)] Success")
	t.Tags("track, service")
	t.Parallel()
	t.WithNewStep("[GetTracksByPartName (service)] Success", func(sCtx provider.StepCtx) {
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
			WithArgs("%boba%", usecase.MinPageSize).
			WillReturnRows(rows)

		var expTracks []*models.TrackMeta
		for _, v := range daoTracks {
			expTracks = append(expTracks, dao.ToModelTrackMeta(v))
		}

		repo := postgres.NewTrackRepository(gormDB)

		tracks, err := usecase.NewTrackUseCase(repo).GetTracksByPartName("boba", -1, -1)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(tracks)
		sCtx.Assert().Equal(expTracks, tracks)
	})
}

func (a *TrackSuite) Test_GetTracksByPartName_Error(t provider.T) {
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

	t.Title("[GetTracksByPartName (service)] Error from db")
	t.Tags("track, service")
	t.Parallel()
	t.WithNewStep("[GetTracksByPartName (service)] Error from db", func(sCtx provider.StepCtx) {
		mock.ExpectQuery("^SELECT (.+) FROM \"tracks\" WHERE name LIKE(.+)$").
			WithArgs("%boba%", usecase.MaxPageSize).
			WillReturnError(assert.AnError)

		repo := postgres.NewTrackRepository(gormDB)

		tracks, err := usecase.NewTrackUseCase(repo).GetTracksByPartName("boba", 0, 1000)

		sCtx.Assert().ErrorIs(err, assert.AnError)
		sCtx.Assert().Nil(tracks)
	})
}
