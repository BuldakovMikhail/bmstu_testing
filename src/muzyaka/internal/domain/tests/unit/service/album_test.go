package service

import (
	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
	mock_repository "src/internal/domain/album/repository/mocks"
	"src/internal/domain/album/usecase"
	"src/internal/lib/testing/builders"
	"src/internal/lib/testing/mother"
	"src/internal/models"
	"testing"
)

type AlbumSuite struct {
	suite.Suite
	t *testing.T
}

func (a *AlbumSuite) Test_GetAlbum_Success(t provider.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	t.Title("[GetAlbum (service)] Success")
	t.Tags("album, service")
	t.Parallel()
	t.WithNewStep("[GetAlbum (service)] Success", func(sCtx provider.StepCtx) {
		respAlbum := builders.AlbumBuilder{}.
			WithId(1).
			WithName("test").
			WithCoverFile([]byte{1, 2, 3}).
			Build()

		repo := mock_repository.NewMockAlbumRepository(c)
		repo.EXPECT().GetAlbum(uint64(1)).Return(respAlbum, nil)

		album, err := usecase.NewAlbumUseCase(repo).GetAlbum(1)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(album)
		sCtx.Assert().Equal(respAlbum, album)
	})
}

func (a *AlbumSuite) Test_GetAlbum_Error(t provider.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	t.Title("[GetAlbum (service)] Error from repository")
	t.Tags("album, service")
	t.Parallel()
	t.WithNewStep("[GetAlbum (service)] Error from repository", func(sCtx provider.StepCtx) {
		repo := mock_repository.NewMockAlbumRepository(c)
		repo.EXPECT().GetAlbum(uint64(1)).Return(nil, assert.AnError)

		album, err := usecase.NewAlbumUseCase(repo).GetAlbum(1)

		sCtx.Assert().ErrorIs(err, assert.AnError)
		sCtx.Assert().Nil(album)
	})
}

func (a *AlbumSuite) Test_GetAllTracks_Success(t provider.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	t.Title("[GetAllTracks (service)] Success")
	t.Tags("album, service")
	t.Parallel()
	t.WithNewStep("[GetAllTracks (service)] Success", func(sCtx provider.StepCtx) {
		var respTracks []*models.TrackMeta
		respTracks = append(respTracks, mother.TrackMetaObjectMother{}.DefaultTrack())

		repo := mock_repository.NewMockAlbumRepository(c)
		repo.EXPECT().GetAllTracks(uint64(1)).Return(respTracks, nil)

		tracks, err := usecase.NewAlbumUseCase(repo).GetAllTracks(1)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(tracks)
		sCtx.Assert().Equal(respTracks, tracks)
	})
}

func (a *AlbumSuite) Test_GetAllTracks_ErrorFromRepo(t provider.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	t.Title("[GetAllTracks (service)] Error from repository")
	t.Tags("album, service")
	t.Parallel()
	t.WithNewStep("[GetAllTracks (service)] Error from repository", func(sCtx provider.StepCtx) {
		repo := mock_repository.NewMockAlbumRepository(c)
		repo.EXPECT().GetAllTracks(uint64(1)).Return(nil, assert.AnError)

		tracks, err := usecase.NewAlbumUseCase(repo).GetAllTracks(1)

		sCtx.Assert().ErrorIs(err, assert.AnError)
		sCtx.Assert().Nil(tracks)
	})
}
