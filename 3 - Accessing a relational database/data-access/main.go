package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//definindo propriedades de conexão
	cfg := mysql.Config{
		User:                 "kevenny",  //os.Getenv(""),
		Passwd:               "MasterOv", //os.Getenv(""),
		Net:                  "tcp",
		AllowNativePasswords: true,
		Addr:                 "localhost:3306",
		DBName:               "recordings",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Conectado!")

	// ================================================ GET BY ARTIST

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album encontrado: %v\n", albums)

	// ================================================ GET ALL ALBUM

	albumsAll, err := albumAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums Encontrados: %v\n", albumsAll)

	// ================================================ ADD ALBUM

	ironMaidenAlbum := Album{
		Title:  "Piece of Mind",
		Artist: "Iron Maiden",
		Price:  99.99,
	}

	albID, err := addAlbum(ironMaidenAlbum)
	if err != nil {
		log.Fatal(err)
	}
	ironMaidenAlbum.ID = albID
	fmt.Printf("ID do album adicionado: %v\n", albID)

	// ================================================ GET BY ID
	//obtendo dados do album que acabei de adicionar
	alb, err := albumByID(albID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album encontrado: %v\n", alb)

	// ================================================ UPDATE ALBUM

	ironMaidenAlbum.Title = "Piece of Mind (2015 Remaster)"
	album, err := updateAlbum(ironMaidenAlbum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album atualizado: %v\n", album)

	// ================================================ DELETE ALBUM

	//deletando o album que acabei de adicionar
	rmStatus, err := removeAlbum(albID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Status da remoção: %v", rmStatus)

}
