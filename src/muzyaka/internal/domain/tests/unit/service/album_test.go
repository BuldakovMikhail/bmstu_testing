package service

import (
	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	mock_repository "src/internal/domain/album/repository/mocks"
	"src/internal/domain/album/usecase"
	"src/internal/lib/testing/builders"
	"testing"
)

type AlbumSuite struct {
	suite.Suite
	t *testing.T
}

func (a *AlbumSuite) Test_GetAlbum(t provider.T) {
	t.Title("[GetAlbum] Success")
	t.Tags("album")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		respAlbum := builders.AlbumBuilder{}.
			WithId(1).
			WithName("test").
			WithCoverFile([]byte{1, 2, 3}).
			BuildModel()

		c := gomock.NewController(t)
		defer c.Finish()

		repo := mock_repository.NewMockAlbumRepository(c)
		repo.EXPECT().GetAlbum(uint64(1)).Return(respAlbum, nil)

		album, err := usecase.NewAlbumUseCase(repo).GetAlbum(1)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(album)
		sCtx.Assert().Equal(respAlbum, album)
	})
}
