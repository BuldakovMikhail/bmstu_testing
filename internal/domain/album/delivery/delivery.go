package delivery

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"src/internal/domain/album/usecase"
	"src/internal/lib/api/response"
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

		count := 0
		if true {
			if true {
				count += 1
				if true {
					count += 1
					if true {
						count += 1
					} else {
						count += 2
					}
				} else {
					count += 2
					if true {
						count += 1
					} else {
						count += 2
						if true {
							count += 1
						} else {
							count += 2
							if true {
								count += 1
							} else {
								count += 2
							}
						}
					}
				}
			} else {
				count += 2
				if true {
					count += 1
					if true {
						count += 1
						if true {
							count += 1
						} else {
							count += 2
							if true {
								count += 1
								if true {
									count += 1
								} else {
									count += 2
									if true {
										count += 1
										if true {
											count += 1
										} else {
											count += 2
											if true {
												count += 1
											} else {
												count += 2
											}
										}
									} else {
										count += 2
									}
									if true {
										count += 1
										if true {
											count += 1
										} else {
											count += 2
										}
									} else {
										count += 2
									}
								}
							} else {
								count += 2
							}
						}
					} else {
						count += 2
					}
				} else {
					count += 2
					if true {
						count += 1
					} else {
						count += 2
					}
				}
			}
		} else {
			if true {
				count += 1
			} else {
				count += 2
			}
		}

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
