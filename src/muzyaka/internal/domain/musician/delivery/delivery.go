package delivery

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"src/internal/domain/musician/usecase"
	"src/internal/lib/api/response"
	"src/internal/models/dto"
	"strconv"
)

// @Summary MusicianGet
// @Security ApiKeyAuth
// @Tags musician
// @Description get musician
// @ID get-musician
// @Accept  json
// @Produce  json
// @Param musician_id   path      int  true  "Musician ID"
// @Success 200 {object} dto.Musician
// @Failure 400,404 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure default {object} response.Response
// @Router /api/musician/{musician_id} [get]
func GetMusician(musicianUseCase usecase.MusicianUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "musician_id")
		aid, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		mus, err := musicianUseCase.GetMusician(aid)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		render.JSON(w, r, dto.ToDtoMusician(mus))
	}
}
