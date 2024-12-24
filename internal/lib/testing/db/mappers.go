package dbhelpers

import (
	"github.com/DATA-DOG/go-sqlmock"
	"src/internal/models/dao"
)

func MapAlbum(album *dao.Album) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "name", "cover_file", "musician_id"}).
		AddRow(album.ID, album.Name, album.Cover, album.MusicianID)
}

func MapTracks(tracks []*dao.Track) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "payload", "name", "album_id"})
	for _, v := range tracks {
		rows.AddRow(v.ID, v.Payload, v.Name, v.AlbumID)
	}

	return rows
}
