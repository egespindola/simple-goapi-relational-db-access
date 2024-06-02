package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/egespindola/simple-goapi-relational-db-access/internal/database"
	"github.com/egespindola/simple-goapi-relational-db-access/internal/service"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "172.25.0.2:3306",
		DBName: "recordings",
	}
	//  mysql native password authentication
	cfg.AllowNativePasswords = true

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connection established!")

	albumDb := database.NewAlbum(db)
	albumService := service.NewAlbumService(*albumDb)

	albums, err := albumService.AlbumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	//
	album, err := albumService.AlbumByID(6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", album)

}
