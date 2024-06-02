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

	return albums, nil

}

// albumById query get album by ID
func (a *AlbumService) AlbumByID(id int64) (database.Album, error) {
	album, err := a.AlbumDB.FindOne(id)

	if err != nil {
		return album, err
	}

	return album, nil

}

// insert an Album into the database
func (a *AlbumService) AlbumAdd(alb database.Album) (int64, error) {
	id, err := a.AlbumDB.InsertAlbum(alb)
	if err != nil {
		return 0, err
	}

	return id, nil
}
