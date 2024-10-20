package e2e

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"net/http"
	"path/filepath"
	"src/cmd/muzyaka"
	dbhelpers "src/internal/lib/testing/db"
	"sync"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Parallel()
	dbMeta, err := dbhelpers.CreateDatabase(context.Background(), filepath.Join("..", "..", "database", "docker-entrypoint-initdb.d", "01-init.sql"))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a test database connection", err)
	}
	defer dbMeta.Terminate(context.Background())

	go muzyaka.App(dbMeta.DB)

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		&E2ESuite{
			client: http.DefaultClient,
		},
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
