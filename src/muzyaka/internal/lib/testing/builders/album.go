package builders

import "src/internal/models"

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

func (a AlbumBuilder) BuildModel() *models.Album {
	return &models.Album{
		Id:        a.Id,
		Name:      a.Name,
		CoverFile: a.CoverFile,
	}
}
