package delivery

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"src/internal/domain/user/usecase"
	"src/internal/lib/api/response"
	"src/internal/models/dto"
	"strconv"
)

// @Summary LikeTrack
// @Security ApiKeyAuth
// @Tags user
// @Description like track
// @ID like-track
// @Accept  json
// @Produce  json
// @Param user_id path int true "user ID"
// @Param input body dto.Like true "liked track"
// @Success 200 {object} response.Response
// @Failure 400,404 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure default {object} response.Response
// @Router /api/user/{user_id}/favorite [post]
func Like(useCase usecase.UserUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "user_id")
		userIDUint, err := strconv.ParseUint(userID, 10, 64)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		var req dto.Like
		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		err = useCase.LikeTrack(userIDUint, req.TrackId)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		render.JSON(w, r, response.OK())
	}
}

// @Summary DislikeTrack
// @Security ApiKeyAuth
// @Tags user
// @Description dislike track
// @ID dislike-track
// @Accept  json
// @Produce  json
// @Param user_id path int true "user ID"
// @Param input body dto.Dislike true "disliked track"
// @Success 200 {object} response.Response
// @Failure 400,404 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure default {object} response.Response
// @Router /api/user/{user_id}/favorite [delete]
func Dislike(useCase usecase.UserUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "user_id")
		userIDUint, err := strconv.ParseUint(userID, 10, 64)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		var req dto.Dislike
		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		err = useCase.DislikeTrack(userIDUint, req.TrackId)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		render.JSON(w, r, response.OK())
	}
}

// @Summary GetAllLiked
// @Security ApiKeyAuth
// @Tags user
// @Description get all liked tracks
// @ID get-all-liked-tracks
// @Accept  json
// @Produce  json
// @Param user_id path int true "user ID"
// @Success 200 {object} dto.TracksMetaCollection
// @Failure 400,404 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure default {object} response.Response
// @Router /api/user/{user_id}/favorite [get]
func GetAllLiked(useCase usecase.UserUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "user_id")
		userIDUint, err := strconv.ParseUint(userID, 10, 64)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		likedTracks, err := useCase.GetAllLikedTracks(userIDUint)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		var dtoLikedTracks []*dto.TrackMeta
		for _, v := range likedTracks {
			dtoLikedTracks = append(dtoLikedTracks, dto.ToDtoTrackMeta(v))
		}

		render.JSON(w, r, dto.TracksMetaCollection{Tracks: dtoLikedTracks})
	}
}
