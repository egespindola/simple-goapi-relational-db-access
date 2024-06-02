package service

import (
	"github.com/egespindola/simple-goapi-relational-db-access/internal/database"
)

type AlbumService struct {
	AlbumDB database.Album
}

func NewAlbumService(albumDB database.Album) *AlbumService {
	return &AlbumService{
		AlbumDB: albumDB,
	}
}

// albumsByArtist queries for albums that have the specified artist name.
func (a *AlbumService) AlbumsByArtist(name string) ([]database.Album, error) {
	albums, err := a.AlbumDB.FindByArtist(name)

	if err != nil {
		return nil, err
	}

	return albums, err

}
