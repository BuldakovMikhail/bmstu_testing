package builders

import (
	"src/internal/models"
	"src/internal/models/dao"
)

type AlbumBuilder struct {
	Id        uint64
	Name      string
	CoverFile []byte
}

func (a AlbumBuilder) WithId(id uint64) AlbumBuilder {
	a.Id = id
	return a
}

func (a AlbumBuilder) WithName(name string) AlbumBuilder {
	a.Name = name
	return a
}

func (a AlbumBuilder) WithCoverFile(coverFile []byte) AlbumBuilder {
	a.CoverFile = coverFile
	return a
}

func (a AlbumBuilder) Build() *models.Album {
	return &models.Album{
		Id:        a.Id,
		Name:      a.Name,
		CoverFile: a.CoverFile,
	}
}

type AlbumDaoBuilder struct {
	Id         uint64
	Name       string
	CoverFile  []byte
	MusicianId uint64
}

func (a AlbumDaoBuilder) WithId(id uint64) AlbumDaoBuilder {
	a.Id = id
	return a
}

func (a AlbumDaoBuilder) WithName(name string) AlbumDaoBuilder {
	a.Name = name
	return a
}

func (a AlbumDaoBuilder) WithCoverFile(coverFile []byte) AlbumDaoBuilder {
	a.CoverFile = coverFile
	return a
}

func (a AlbumDaoBuilder) WithMusicianId(id uint64) AlbumDaoBuilder {
	a.MusicianId = id
	return a
}

func (a AlbumDaoBuilder) Build() *dao.Album {
	return &dao.Album{
		ID:         a.Id,
		Name:       a.Name,
		Cover:      a.CoverFile,
		MusicianID: a.MusicianId,
	}
}
