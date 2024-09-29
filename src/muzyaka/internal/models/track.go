package models

type TrackMeta struct {
	Id    uint64
	Name  string
	Genre string
}

type TrackObject struct {
	Payload []byte
}

func (t *TrackObject) ExtractMeta() *TrackMeta {
	return &t.TrackMeta
}
