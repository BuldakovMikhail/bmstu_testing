package e2e

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"net/http"
	"sync"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Parallel()

	db, ids, err := InitDatabase(context.Background())
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a test database connection", err)
	}
	defer ClearTestDB(db, ids)

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
