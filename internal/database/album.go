package database

import (
	"database/sql"
	"fmt"
)

type Album struct {
	db     *sql.DB
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func NewAlbum(db *sql.DB) *Album {
	return &Album{db: db}
}

func (a *Album) FindByArtist(name string) ([]Album, error) {
	var albums []Album
	rows, err := a.db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v ", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

func (a *Album) FindOne(id int64) (Album, error) {
	var alb Album

	row := a.db.QueryRow("SELECT * FROM album WHERE ID = ?", id)

	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			// return alb, fmt.Errorf("albumsById %d: no such album", id)
			return alb, err
		}
		// return alb, fmt.Errorf("albumsById %d: %v", id, err)
		return alb, err

	}

	return alb, nil

}

func (a *Album) InsertAlbum(alb Album) (int64, error) {
	result, err := a.db.Exec("INSERT INTO album (title, artist, price) VALUES (?,?,?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
