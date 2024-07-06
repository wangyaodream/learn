package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type album struct {
    ID string `json:"id"`
    Title string `json:"title"`
    Artist string `json:"artist"`
    Price float64 `json:"price"`
}


var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarach Vaughan and Clifford Brown", Artist: "Sarach Vaughan", Price: 39.99},
}


func getAlbums(c *gin.Context) {
    // 将结构体序列化为JSON
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
    var newAlbum album
    
    // 通过BindJSON去直接绑定request中的body，去匹配album类型的newAlbum
    // 当request的body中没有所匹配的数据时，将会以0值的形式进行添加
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
    // 这里的参数和router中设定的参数名一致，用来捕获参数值
    id := c.Param("foo")

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
    router.GET("/albums/:foo", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}
