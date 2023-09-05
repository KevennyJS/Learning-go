package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.PUT("/albums", updateAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.Run(":6666")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(context *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(context *gin.Context) {
	id := context.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range albums {
		if album.ID == id {
			context.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album não encontrado"})
}

func updateAlbums(context *gin.Context) {
	var editedAlbum Album
	if err := context.BindJSON(&editedAlbum); err != nil {
		return
	}

	for index, album := range albums {
		if album.ID == editedAlbum.ID {
			albums[index] = editedAlbum
			context.IndentedJSON(http.StatusOK, editedAlbum)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album não encontrado"})
}

func deleteAlbumByID(context *gin.Context) {
	id := context.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			context.JSON(http.StatusOK, gin.H{"message": "Album removido"})
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album não encontrado"})
}
