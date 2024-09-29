package dbhelpers

import (
	"github.com/DATA-DOG/go-sqlmock"
	"src/internal/models/dao"
)

func MapAlbum(album *dao.Album) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "name", "cover_file", "musician_id"}).
		AddRow(album.ID, album.Name, album.Cover, album.MusicianID)
}
