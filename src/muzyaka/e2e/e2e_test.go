package e2e

import (
	"net/http"
	url2 "net/url"
	"src/internal/models/dto"
	"strconv"

	"github.com/go-chi/render"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type E2ESuite struct {
	suite.Suite

	client *http.Client
}

func (s *E2ESuite) Test_SearchTrack(t provider.T) {
	t.Title("[SearchTrack (e2e)] Success")
	t.Tags("track, e2e, SearchTrack")
	t.Parallel()
	t.WithNewStep("[SearchTrack (e2e)] Success", func(sCtx provider.StepCtx) {
		request, err := http.NewRequest("GET", "http://localhost:8080/api/track", nil)
		if err != nil {
			t.Fatalf("Unable to create request: %s", err.Error())
		}
		request.URL.RawQuery = url2.Values{
			"q":         {"test"},
			"page":      {strconv.Itoa(-1)},
			"page_size": {strconv.Itoa(-1)},
		}.Encode()
		request.Close = true
		resp, err := s.client.Do(request)
		if err != nil {
			t.Fatalf("Error during request: %s", err.Error())
		}

		defer resp.Body.Close()

		var decodedResp dto.TracksMetaCollection
		err = render.DecodeJSON(resp.Body, &decodedResp)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(1, len(decodedResp.Tracks))
		sCtx.Assert().Equal("test", decodedResp.Tracks[0].Name)
	})
}
