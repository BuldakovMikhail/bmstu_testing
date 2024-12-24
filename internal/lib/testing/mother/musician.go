package mother

import "src/internal/models/dao"

type MusicianDaoMother struct{}

func (m MusicianDaoMother) DefaultMusician() *dao.Musician {
	return &dao.Musician{
		ID:   0,
		Name: "Prorok Sunboy",
	}
}
