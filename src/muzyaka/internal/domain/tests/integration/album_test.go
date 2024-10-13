package integration

import (
	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"src/internal/domain/album/repository/postgres"
	"src/internal/lib/testing/builders"
	dbhelpers "src/internal/lib/testing/db"
	"src/internal/models"
)

type AlbumIntegrationSuite struct {
	suite.Suite

	TestDB *dbhelpers.TestDatabaseMeta
}

func (a *AlbumIntegrationSuite) Test_GetAlbum_Success(t provider.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	t.Title("[GetAlbum (integration)] Success")
	t.Tags("album, integration")
	t.Parallel()
	t.WithNewStep("[GetAlbum (integration)] Success", func(sCtx provider.StepCtx) {
		album := builders.AlbumBuilder{}.
			WithId(a.TestDB.IDs["albumId"]).
			WithName("test").
			WithCoverFile([]byte{1, 2, 3}).
			Build()
		repo := postgres.NewAlbumRepository(a.TestDB.DB)

		respAlbum, err := repo.GetAlbum(a.TestDB.IDs["albumId"])

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(respAlbum)
		sCtx.Assert().Equal(album, respAlbum)
	})
}

func (a *AlbumIntegrationSuite) Test_GetAlbum_NotFound(t provider.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	t.Title("[GetAlbum (integration)] Not Found")
	t.Tags("album, integration")
	t.Parallel()
	t.WithNewStep("[GetAlbum (integration)] Not Found", func(sCtx provider.StepCtx) {
		repo := postgres.NewAlbumRepository(a.TestDB.DB)

		respAlbum, err := repo.GetAlbum(a.TestDB.IDs["albumId"] + 1)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(respAlbum)
	})
}

func (a *AlbumIntegrationSuite) Test_GetAllTracks_Success(t provider.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	t.Title("[GetAllTracks (integration)] Success")
	t.Tags("album, integration")
	t.Parallel()
	t.WithNewStep("[GetAllTracks (integration)] Success", func(sCtx provider.StepCtx) {
		track := builders.TrackMetaBuilder{}.
			WithId(a.TestDB.IDs["trackId"]).
			WithName("test").
			Build()
		repo := postgres.NewAlbumRepository(a.TestDB.DB)

		respTracks, err := repo.GetAllTracks(a.TestDB.IDs["albumId"])

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(respTracks)
		sCtx.Assert().Equal([]*models.TrackMeta{track}, respTracks)
	})
}

func (a *AlbumIntegrationSuite) Test_GetAllTracks_NotFound(t provider.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	t.Title("[GetAllTracks (integration)] Not Found")
	t.Tags("album, integration")
	t.Parallel()
	t.WithNewStep("[GetAllTracks (integration)] Not Found", func(sCtx provider.StepCtx) {
		repo := postgres.NewAlbumRepository(a.TestDB.DB)

		respTracks, err := repo.GetAllTracks(a.TestDB.IDs["albumId"] + 1)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(respTracks)
	})
}
