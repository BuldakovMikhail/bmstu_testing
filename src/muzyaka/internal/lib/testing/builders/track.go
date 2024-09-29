package builders

import "src/internal/models"

type TrackMetaBuilder struct {
	Id   uint64
	Name string
}

func (t TrackMetaBuilder) WithId(id uint64) TrackMetaBuilder {
	t.Id = id
	return t
}

func (t TrackMetaBuilder) WithName(name string) TrackMetaBuilder {
	t.Name = name
	return t
}

func (t TrackMetaBuilder) BuildTrackMeta() *models.TrackMeta {
	return &models.TrackMeta{
		Id:   t.Id,
		Name: t.Name,
	}
}
