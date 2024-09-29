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
	"src/internal/models/dao"
	"testing"
)

type AlbumSuite struct {
	suite.Suite
	t *testing.T
}

func (a *AlbumSuite) Test_GetAlbum_Success(t provider.T) {
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

	t.Title("[GetAlbum] Success")
	t.Tags("album")
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

func (a *AlbumSuite) Test_GetAlbum_Error(t provider.T) {
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

	t.Title("[GetAlbum] Success")
	t.Tags("album")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		mock.ExpectQuery("^SELECT (.+) FROM \"albums\" WHERE id = (.+)$").
			WithArgs(1, 1).
			WillReturnError(assert.AnError)

		album, err := postgres.NewAlbumRepository(gormDB).GetAlbum(1)

		sCtx.Assert().ErrorIs(err, assert.AnError)
		sCtx.Assert().Nil(album)
	})
}
