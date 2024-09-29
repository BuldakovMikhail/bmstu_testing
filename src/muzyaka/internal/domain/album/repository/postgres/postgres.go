package postgres

import (
	"gorm.io/gorm"
	"src/internal/domain/album/repository"
)

type albumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) repository.AlbumRepository {
	return &albumRepository{db: db}
}
