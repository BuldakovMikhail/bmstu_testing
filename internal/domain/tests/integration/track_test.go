package integration

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"src/internal/domain/track/repository/postgres"
	"src/internal/domain/track/usecase"
	"src/internal/lib/testing/builders"
	dbhelpers "src/internal/lib/testing/db"
	"src/internal/models"
)

type TrackIntegrationSuite struct {
	suite.Suite

	TestDB *dbhelpers.TestDatabaseMeta
}

func (ts *TrackIntegrationSuite) Test_GetTrack_Success(t provider.T) {
	t.Title("[GetTrack (integration)] Success")
	t.Tags("track, integration")
	t.Parallel()
	t.WithNewStep("[GetTrack (integration)] Success", func(sCtx provider.StepCtx) {
		track := builders.TrackObjectBuilder{}.
			WithPayload([]byte{4, 5, 6}).
			Build()
		repo := postgres.NewTrackRepository(ts.TestDB.DB)

		respTrack, err := usecase.NewTrackUseCase(repo).GetTrack(ts.TestDB.IDs["trackId"])

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(respTrack)
		sCtx.Assert().Equal(track, respTrack)
	})
}

func (ts *TrackIntegrationSuite) Test_GetTrack_NotFound(t provider.T) {
	t.Title("[GetTrack (integration)] Not Found")
	t.Tags("track, integration")
	t.Parallel()
	t.WithNewStep("[GetTrack (integration)] Not Found", func(sCtx provider.StepCtx) {
		repo := postgres.NewTrackRepository(ts.TestDB.DB)

		respTrack, err := usecase.NewTrackUseCase(repo).GetTrack(ts.TestDB.IDs["trackId"] + 1)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(respTrack)
	})
}

func (ts *TrackIntegrationSuite) Test_GetTracksByPartName_Success(t provider.T) {
	t.Title("[GetTracksByPartName (integration)] Success")
	t.Tags("track, integration")
	t.Parallel()
	t.WithNewStep("[GetTracksByPartName (integration)] Success", func(sCtx provider.StepCtx) {
		track := builders.TrackMetaBuilder{}.
			WithId(ts.TestDB.IDs["trackId"]).
			WithName("test").
			Build()
		repo := postgres.NewTrackRepository(ts.TestDB.DB)

		respTracks, err := usecase.NewTrackUseCase(repo).GetTracksByPartName("test", 1, -1)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(respTracks)
		sCtx.Assert().Equal([]*models.TrackMeta{track}, respTracks)
	})
}

func (ts *TrackIntegrationSuite) Test_GetTracksByPartName_NotFound(t provider.T) {
	t.Title("[GetTracksByPartName (integration)] Not Found")
	t.Tags("track, integration")
	t.Parallel()
	t.WithNewStep("[GetTracksByPartName (integration)] Not Found", func(sCtx provider.StepCtx) {
		repo := postgres.NewTrackRepository(ts.TestDB.DB)

		respTracks, err := usecase.NewTrackUseCase(repo).GetTracksByPartName("asd", 1, -1)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Nil(respTracks)
	})
}
