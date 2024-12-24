package builders

import (
	"src/internal/models"
	"src/internal/models/dao"
)

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

func (t TrackMetaBuilder) Build() *models.TrackMeta {
	return &models.TrackMeta{
		Id:   t.Id,
		Name: t.Name,
	}
}

type TrackDaoMetaBuilder struct {
	ID      uint64
	Payload []byte
	Name    string
	AlbumID uint64
}

func (t TrackDaoMetaBuilder) WithId(id uint64) TrackDaoMetaBuilder {
	t.ID = id
	return t
}

func (t TrackDaoMetaBuilder) WithName(name string) TrackDaoMetaBuilder {
	t.Name = name
	return t
}

func (t TrackDaoMetaBuilder) WithPayload(payload []byte) TrackDaoMetaBuilder {
	t.Payload = payload
	return t
}

func (t TrackDaoMetaBuilder) WithAlbumId(id uint64) TrackDaoMetaBuilder {
	t.AlbumID = id
	return t
}

func (t TrackDaoMetaBuilder) Build() *dao.Track {
	return &dao.Track{
		ID:      t.ID,
		Payload: t.Payload,
		Name:    t.Name,
		AlbumID: t.AlbumID,
	}
}

type TrackObjectBuilder struct {
	Payload []byte
}

func (t TrackObjectBuilder) WithPayload(payload []byte) TrackObjectBuilder {
	t.Payload = payload
	return t
}

func (t TrackObjectBuilder) Build() *models.TrackObject {
	return &models.TrackObject{
		Payload: t.Payload,
	}
}
