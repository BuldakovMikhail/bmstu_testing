package dto

import "src/internal/models"

type TrackMeta struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type TrackMetaWithoutId struct {
	Name string `json:"name"`
}

type TrackObjectWithoutId struct {
	TrackMetaWithoutId
	Payload []byte `json:"payload"`
}

type TrackObject struct {
	TrackMeta
	Payload []byte `json:"payload"`
}

type TrackObjectWithSource struct {
	TrackMeta
	Payload []byte `json:"payload"`
}

type TracksMetaCollection struct {
	Tracks []*TrackMeta `json:"tracks"`
}

func ToDtoTrackMeta(m *models.TrackMeta) *TrackMeta {
	return &TrackMeta{
		Id:   m.Id,
		Name: m.Name,
	}
}

func ToDtoTrackObjectWithSource(t *models.TrackObject) *TrackObjectWithSource {
	return &TrackObjectWithSource{
		Payload: t.Payload,
	}
}
