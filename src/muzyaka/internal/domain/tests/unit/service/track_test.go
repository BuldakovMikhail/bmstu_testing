package service

//func (u *usecase) GetTracksByPartName(name string, page int, pageSize int) ([]*models.TrackMeta, error) {
//	if page <= 0 {
//		page = 1
//	}
//
//	switch {
//	case pageSize > MaxPageSize:
//		pageSize = MaxPageSize
//	case pageSize < MinPageSize:
//		pageSize = MinPageSize
//	}
//
//	offset := (page - 1) * pageSize
//	tracks, err := u.trackRep.GetTracksByPartName(name, offset, pageSize)
//	if err != nil {
//		return nil, errors.Wrap(err, "track.usecase.GetTracksByPartName error while get")
//	}
//
//	return tracks, nil
//}
//
//func (u *usecase) GetTrack(id uint64) (*models.TrackObject, error) {
//	res, err := u.trackRep.GetTrack(id)
//	if err != nil {
//		return nil, errors.Wrap(err, "track.usecase.GetTrack error while get")
//	}
//
//	return res, nil
//}
