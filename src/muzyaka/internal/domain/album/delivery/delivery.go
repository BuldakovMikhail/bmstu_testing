package delivery

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"src/internal/domain/album/usecase"
	"src/internal/lib/api/response"
	"src/internal/models"
	"src/internal/models/dto"
	"strconv"
)

// @Summary GetAlbum
// @Security ApiKeyAuth
// @Tags album
// @Description get album by ID
// @ID get-album
// @Accept  json
// @Produce  json
// @Param id path int true "album ID"
// @Success 200 {object} dto.Album
// @Failure 400,404,405 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure default {object} response.Response
// @Router /api/album/{id} [get]
func GetAlbum(useCase usecase.AlbumUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		albumID := chi.URLParam(r, "id")
		albumIDUint, err := strconv.ParseUint(albumID, 10, 64)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		album, err := useCase.GetAlbum(albumIDUint)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		render.JSON(w, r, dto.ToDtoAlbum(album))
	}
}

// @Summary DeleteAlbum
// @Security ApiKeyAuth
// @Tags album
// @Description delete album
// @ID delete-album
// @Accept  json
// @Produce  json
// @Param id path int true "album ID"
// @Success 200 {object} response.Response
// @Failure 400,404,405 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure default {object} response.Response
// @Router /api/album/{id} [delete]
func DeleteAlbum(useCase usecase.AlbumUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		albumID := chi.URLParam(r, "id")
		albumIDUint, err := strconv.ParseUint(albumID, 10, 64)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		err = useCase.DeleteAlbum(albumIDUint)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		render.JSON(w, r, response.OK())
	}
}

// @Summary AddAlbumWithTracks
// @Security ApiKeyAuth
// @Tags musician
// @Description add album with tracks
// @ID add-album-with-tracks
// @Accept  json
// @Produce  json
// @Param musician_id path int true "musician ID"
// @Param input body dto.AlbumWithTracks true "album info"
// @Success 200 {object} dto.CreateAlbumResponse
// @Failure 400,404,405 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure default {object} response.Response
// @Router /api/musician/{musician_id}/album [post]
func AddAlbumWithTracks(useCase usecase.AlbumUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		musicianID := chi.URLParam(r, "musician_id")
		musicianIDUint, err := strconv.ParseUint(musicianID, 10, 64)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		var req dto.AlbumWithTracks
		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		var modelTracks []*models.TrackObject
		for _, v := range req.Tracks {
			modelTracks = append(modelTracks, dto.ToModelTrackObjectWithoutId(v, 0, ""))
		}

		albumID, err := useCase.AddAlbumWithTracks(dto.ToModelAlbumWithId(0, &req.AlbumWithoutId), modelTracks, musicianIDUint)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		render.JSON(w, r, dto.CreateAlbumResponse{Id: albumID})
	}
}

// @Summary GetAllTracks
// @Security ApiKeyAuth
// @Tags album
// @Description get all tracks from album
// @ID get-all-tracks-from-album
// @Accept  json
// @Produce  json
// @Param id path int true "album ID"
// @Success 200 {object} dto.TracksMetaCollection
// @Failure 400,404,405 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure default {object} response.Response
// @Router /api/album/{id}/tracks [get]
func GetAllTracks(useCase usecase.AlbumUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		albumID := chi.URLParam(r, "id")
		albumIDUint, err := strconv.ParseUint(albumID, 10, 64)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		tracks, err := useCase.GetAllTracks(albumIDUint)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		var res []*dto.TrackMeta
		for _, v := range tracks {
			res = append(res, dto.ToDtoTrackMeta(v))
		}

		render.JSON(w, r, dto.TracksMetaCollection{Tracks: res})
	}
}
