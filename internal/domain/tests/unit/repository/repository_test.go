package repository

import (
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"sync"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Parallel()

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		&AlbumRepoSuite{t: t},
		&TrackRepoSuite{t: t},
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
