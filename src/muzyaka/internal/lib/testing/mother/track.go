package mother

import "src/internal/models"

type TrackMetaObjectMother struct {
}

func (t TrackMetaObjectMother) DefaultTrack() *models.TrackMeta {
	return &models.TrackMeta{
		Id:   1,
		Name: "default",
	}
}
