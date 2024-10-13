package integration

import (
	"github.com/ozontech/allure-go/pkg/framework/suite"
	dbhelpers "src/internal/lib/testing/db"
)

type TrackIntegrationSuite struct {
	suite.Suite

	TestDB *dbhelpers.TestDatabaseMeta
}
