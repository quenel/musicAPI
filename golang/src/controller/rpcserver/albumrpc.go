package rpcserver

import(
	"net/http"
	"log"
	"fmt"
	"github.com/labstack/echo"
)

const (
	GetAlbumsIdsRPCName = "getAlbumsIds"
	GetAlbumRPCName = "getAlbum"
	GetAlbumCloudWordRPCName = "getAlbumCloudWord"
)

func GetAlbumsIdsRPC(c echo.Context) error {
	return c.JSON(http.StatusOK, Server.MusicRepo.GetAlbumsIds())
}

func GetAlbumRPC(c echo.Context) error {
	query := struct{
		Id int `json:"id"`	
	}{}
	if err := c.Bind(&query) ; err != nil {
		log.Println("error binding request, error:", err)
		return c.JSON(http.StatusBadRequest, struct{}{})
	}

	album, exist := Server.MusicRepo.GetAlbum(query.Id)
	if !exist {
		log.Println("error :", fmt.Errorf("album id %d does not exists", query.Id))
		return c.JSON(http.StatusNotFound, struct{}{})
	}

	res := map[string]interface{}{
		"id" : album.Id,
		"name" : album.Name,
		"artist" : map[string]interface{}{
			"id" : album.Artist.Id,
		},
		"tracks" : album.GetTracksIds(),
	}

	return c.JSON(http.StatusOK, res)
}

func GetAlbumCloudWordRPC(c echo.Context) error {
	query := struct {
		Id 		int `json:"id"`
		Size 	int `json:"size"`
	}{}
	if err := c.Bind(&query) ; err != nil {
		log.Println("error binding request, error:", err)
		return c.JSON(http.StatusBadRequest, struct{}{})
	}
	
	album, exist := Server.MusicRepo.GetAlbum(query.Id)
	if !exist {
		log.Println("error :", fmt.Errorf("album id %d does not exists", query.Id))
		return c.JSON(http.StatusNotFound, struct{}{})
	}
	if query.Size < 0 {
		log.Println("cannot build cloud of size: %d", query.Size)
		return c.JSON(http.StatusBadRequest, struct{}{})
	}
	return c.JSON(http.StatusOK, album.GetCloudWord(query.Size))
}