package integration

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"path/filepath"
	dbhelpers "src/internal/lib/testing/db"
	"sync"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Parallel()

	dbMeta, err := dbhelpers.CreateDatabase(context.Background(), filepath.Join("..", "..", "..", "..", "..", "database", "docker-entrypoint-initdb.d", "01-init.sql"))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a test database connection", err)
	}
	defer dbMeta.Terminate(context.Background())

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		&AlbumIntegrationSuite{TestDB: dbMeta},
		&TrackIntegrationSuite{TestDB: dbMeta},
	}
	wg.Add(len(suits))

	for _, s := range suits {
		go func() {
			suite.RunSuite(t, s)
			wg.Done()
		}()
	}

	wg.Wait()
}
