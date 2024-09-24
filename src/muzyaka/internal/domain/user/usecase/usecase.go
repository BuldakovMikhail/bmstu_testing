package usecase

import (
	"github.com/pkg/errors"
	usecase2 "src/internal/domain/auth/usecase"
	repository2 "src/internal/domain/track/repository"
	"src/internal/domain/user/repository"
	"src/internal/models"
)

type UserUseCase interface {
	LikeTrack(userId uint64, trackId uint64) error
	DislikeTrack(userId uint64, trackId uint64) error
	GetAllLikedTracks(userId uint64) ([]*models.TrackMeta, error)
}

type usecase struct {
	userRep   repository.UserRepository
	trackRep  repository2.TrackRepository
	encryptor usecase2.Encryptor
}

func NewUserUseCase(rep repository.UserRepository, trackRep repository2.TrackRepository, encryptor usecase2.Encryptor) UserUseCase {
	return &usecase{userRep: rep, trackRep: trackRep, encryptor: encryptor}
}

func (u *usecase) GetAllLikedTracks(userId uint64) ([]*models.TrackMeta, error) {
	trackIds, err := u.userRep.GetAllLikedTracks(userId)
	if err != nil {
		return nil, errors.Wrap(err, "user.usecase.GetAllLikedTracks error while get")
	}

	var trackMeta []*models.TrackMeta
	for _, v := range trackIds {
		track, err := u.trackRep.GetTrack(v)
		if err != nil {
			return nil, errors.Wrap(err, "user.usecase.GetAllLikedTracks error while get")
		}

		trackMeta = append(trackMeta, track)
	}

	return trackMeta, nil
}

func (u *usecase) LikeTrack(userId uint64, trackId uint64) error {
	err := u.userRep.LikeTrack(userId, trackId)

	if err != nil {
		return errors.Wrap(err, "user.usecase.LikeTrack error while add")
	}

	return nil
}

func (u *usecase) DislikeTrack(userId uint64, trackId uint64) error {
	err := u.userRep.DislikeTrack(userId, trackId)

	if err != nil {
		return errors.Wrap(err, "user.usecase.DislikeTrack error while delete")
	}

	return nil
}
