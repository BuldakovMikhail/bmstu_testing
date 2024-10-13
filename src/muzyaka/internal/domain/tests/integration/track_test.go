package integration

import (
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"src/internal/domain/track/usecase"
)

type TrackIntegrationSuite struct {
	suite.Suite

	service usecase.TrackUseCase
}
