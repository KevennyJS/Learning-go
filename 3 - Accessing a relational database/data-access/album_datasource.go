package main

import (
	"database/sql"
	"errors"
	"fmt"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// albumsByArtist query para alguns de um artista especifico
func albumsByArtist(name string) ([]Album, error) {
	// O Album slice para guardar os dados retornados
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Faça um loop pelas linhas, usando Scan para atribuir dados de coluna a campos struct.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumAll retorna todos os albums
func albumAll() ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("albumAll: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsAll: %v", err)
		}
		albums = append(albums, alb)
	}

	return albums, nil
}

// albumByID consultas para o álbum com o ID especificado.
func albumByID(id int64) (Album, error) {
	// Um álbum para armazenar dados da linha retornada.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adiciona um album especifico na base de dados
// retorna o ID da nova entidade
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

// removeAlbum remove um album especifico pelo id
// retorna o status da ação
func removeAlbum(id int64) (string, error) {
	_, err := db.Exec("DELETE FROM album WHERE Id = ?", id)
	if err != nil {
		return "Falhou!", fmt.Errorf("removeAlbum: %v", err)
	}
	return "Sucesso!", nil
}

// updateAlbum atualiza um album com os valores recebidos por parametro
// retorna o album atualizado
func updateAlbum(alb Album) (Album, error) {
	if alb.ID == 0 {
		return alb, fmt.Errorf("dados invalidos")
	}

	_, err := db.Exec("UPDATE album set title=?, artist=?, price=? WHERE id = ?", alb.Title, alb.Artist, alb.Price, alb.ID)
	if err != nil {
		return alb, fmt.Errorf("updateAlbum: %v", err)
	}
	return alb, nil
}
