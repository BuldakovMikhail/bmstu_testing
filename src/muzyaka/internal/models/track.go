package models

type TrackMeta struct {
	Id    uint64
	Name  string
	Genre string
}

type TrackObject struct {
	Payload []byte
}
