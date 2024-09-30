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
	"src/internal/models/dao"
	"testing"
)

//func (u *usecase) GetTracksByPartName(name string, page int, pageSize int) ([]*models.TrackMeta, error) {
//	if page <= 0 {
//		page = 1
//	}
//
//	switch {
//	case pageSize > MaxPageSize:
//		pageSize = MaxPageSize
//	case pageSize < MinPageSize:
//		pageSize = MinPageSize
//	}
//
//	offset := (page - 1) * pageSize
//	tracks, err := u.trackRep.GetTracksByPartName(name, offset, pageSize)
//	if err != nil {
//		return nil, errors.Wrap(err, "track.usecase.GetTracksByPartName error while get")
//	}
//
//	return tracks, nil
//}
//
//func (u *usecase) GetTrack(id uint64) (*models.TrackObject, error) {
//	res, err := u.trackRep.GetTrack(id)
//	if err != nil {
//		return nil, errors.Wrap(err, "track.usecase.GetTrack error while get")
//	}
//
//	return res, nil
//}

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

	t.Title("[GetTrack] Success")
	t.Tags("track")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
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

	t.Title("[GetTrack] Error from db")
	t.Tags("track")
	t.Parallel()
	t.WithNewStep("Error from db", func(sCtx provider.StepCtx) {
		mock.ExpectQuery("^SELECT (.+) FROM \"tracks\" WHERE id = (.+)$").
			WithArgs(1, 1).
			WillReturnError(assert.AnError)

		repo := postgres.NewTrackRepository(gormDB)

		track, err := usecase.NewTrackUseCase(repo).GetTrack(1)

		sCtx.Assert().ErrorIs(err, assert.AnError)
		sCtx.Assert().Nil(track)
	})
}
