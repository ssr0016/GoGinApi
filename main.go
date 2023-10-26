package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Developing a RESTful API with Go and Gin

//To keep things simple for the tutorial, you’ll store data in memory. A more typical API would interact with a database.

// Album struct. You’ll use this to store album data in memory.
//Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are serialized into JSON. Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON.

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//Slice of album structs containing data you’ll use to start.

// albums slice to seed record data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Write a handler return all items
// When the client makes a request at GET /albums, you want to return all the albums as JSON.
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Write a handler to add a new item
// When the client makes a POST request at /albums, you want to add the album described in the request body to the existing albums’ data.
// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album
	//Call BindJSON to bind the recieved JSON to
	//newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//Write a handler to return a specific item
//When the client makes a request to GET /albums/[id], you want to return the album whose ID matches the id path parameter.

// getAlbumID locates the album whose ID value matches id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//Loop over the list of albums, looking for
	//an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")

}
